import json
from datetime import datetime, timezone

USERS_FILE = "accounts.json"
AUDIT_LOG_FILE = "audit.log"

SUCCESS = "success"
FAILED = "failed"


# ---------- File I/O ----------
def load_users():
    """Load users from JSON file"""
    with open(USERS_FILE) as f:
        return json.load(f)


def save_users(users):
    """Save users to JSON file"""
    with open(USERS_FILE, "w") as f:
        json.dump(users, f, indent=2)


# ---------- Search ----------
def find_user_by_id(users, user_id):
    """Find a user by ID"""
    for user in users:
        if user["id"] == user_id:
            return user
    return {"ERROR": FAILED}


def find_user_by_name(users, username):
    """Find a user by username"""
    for user in users:
        if user["username"] == username:
            return user
    return {"ERROR": FAILED}


def find_user_by_email(users, email):
    """Find a user by email"""
    for user in users:
        if user["email"] == email:
            return user
    return {"ERROR": FAILED}


# ---------- API Functions ----------
def login(email, password):
    """Check email and password and update last_login"""
    users = load_users()
    user = find_user_by_email(users, email)

    if user is None:
        return {"STATUS": "user not found"}

    if not user["is_active"]:
        return {"STATUS": "user inactive"}

    if user["password"] != password:
        return {"STATUS": "wrong password"}

    user["last_login"] = datetime.now(timezone.utc).isoformat()
    save_users(users)

    return {"STATUS": SUCCESS}


def add_user(username, email, password):
    """Add a new user"""
    users = load_users()

    if find_user_by_email(users, email) is not None:
        return "user already exists"

    if users:
        new_id = max(user["id"] for user in users) + 1
    else:
        new_id = 0

    users.append(
        {
            "id": new_id,
            "username": username,
            "email": email,
            "password": password,
            "is_active": True,
            "last_login": None,
        }
    )

    save_users(users)
    users[new_id].pop("password")
    return users[new_id]
    


def delete_user(email):
    """Delete a user permanently"""
    users = load_users()
    user = find_user_by_email(users, email)

    if user is None:
        return {"ERROR": "user not found"}

    users.remove(user)
    save_users(users)

    log_action("DELETE_USER", email)
    return user


def deactivate_user(email):
    """Deactivate a user account"""
    users = load_users()
    user = find_user_by_email(users, email)

    if user is None:
        return {"ERROR": "user not found"}

    user["is_active"] = False
    save_users(users)
    return user


def list_users():
    """Print all users"""
    users = load_users()
    return [u.pop("password", None) or u for u in users]

def retrieve_user(email):
    """Retrieve user details by email"""
    users = load_users()
    user = find_user_by_email(users, email)
    user.pop("password", None)  # Remove password for security

    if user is None:
        return {"ERROR": "user not found"}

    return user

def change_email(email, new_email):
    """Change user's email address"""
    users = load_users()
    user = find_user_by_email(users, email)

    if user is None:
        return {"ERROR": "user not found"}

    user["email"] = new_email
    save_users(users)
    return user


def change_password(email, old_password, new_password):
    """Change user's password"""
    users = load_users()
    user = find_user_by_email(users, email)

    if user is None:
        return {"ERROR": "user not found"}

    if user["password"] != old_password:
        return {"ERROR": "wrong password"}

    user["password"] = new_password
    save_users(users)

    log_action("CHANGE_PASSWORD", email)
    return {"status": SUCCESS}


# ---------- Audit Logging ----------
def log_action(action, email):
    """Write an audit log"""
    timestamp = datetime.now(timezone.utc).isoformat()
    with open(AUDIT_LOG_FILE, "a") as f:
        f.write(f"{timestamp} {action} {email}\n")


if __name__ == "__main__":
    list_users()
