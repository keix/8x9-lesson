import json
from datetime import datetime

USERS_FILE = "accounts.json"


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


def find_user(users, username):
    """Find a user by username"""
    for user in users:
        if user["username"] == username:
            return user
    return None


# ---------- Operations ----------


def login(username, password):
    """Check username and password and update last_login"""
    users = load_users()
    user = find_user(users, username)

    if user is None:
        return "user not found"

    if not user["is_active"]:
        return "account disabled"

    if user["password"] != password:
        return "wrong password"

    user["last_login"] = datetime.utcnow().isoformat() + "Z"
    save_users(users)

    return "login ok"


def list_users():
    """Print all users"""
    users = load_users()
    for user in users:
        status = "active" if user["is_active"] else "inactive"
        print(user["id"], user["username"], status)


def add_user(username, email, password):
    """Add a new user"""
    users = load_users()

    if find_user(users, username) is not None:
        return "username already exists"

    new_id = max(user["id"] for user in users) + 1

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
    return "user added"


def deactivate_user(username):
    """Deactivate a user account"""
    users = load_users()
    user = find_user(users, username)

    if user is None:
        return "user not found"

    user["is_active"] = False
    save_users(users)
    return "user deactivated"


def change_email(username, new_email):
    """Change user's email address"""
    users = load_users()
    user = find_user(users, username)

    if user is None:
        return "user not found"

    user["email"] = new_email
    save_users(users)
    return "email updated"


# ---------- Example usage ----------

if __name__ == "__main__":
    print(login("madoka", "madoka123"))
    print(add_user("kyube", "kyube@example.com", "kyube123"))
    print(deactivate_user("mami"))
    list_users()
