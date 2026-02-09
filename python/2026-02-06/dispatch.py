import json
import sys
import users


def dispatch(request):
    """Dispatch JSON request to appropriate function"""
    command = request.get("command")

    if command == "login":
        return users.login(request["email"], request["password"])

    elif command == "add":
        return users.add_user(
            request["username"], request["email"], request["password"]
        )

    elif command == "delete":
        return users.delete_user(request["email"])

    elif command == "deactivate":
        return users.deactivate_user(request["email"])

    elif command == "list":
        return users.list_users()

    elif command == "retrieve":
        return users.retrieve_user(request["email"])

    elif command == "change-email":
        return users.change_email(request["email"], request["new_email"])

    elif command == "change-password":
        return users.change_password(
            request["email"], request["old_password"], request["new_password"]
        )

    else:
        return f"unknown command: {command}"


if __name__ == "__main__":
    request = json.loads(sys.argv[1])
    result = dispatch(request)
    print(json.dumps(result))
