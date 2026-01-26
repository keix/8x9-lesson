import requests


def login(username, password):
    url = "http://localhost:18080/login"
    data = {"username": username, "password": password}

    try:
        response = requests.post(url, json=data)
        if response.status_code == 200:
            result = response.json()
            print(f"Login successful! Token: {result['token']}")
            return result["token"]
        else:
            print(f"Login failed: {response.json()}")
            return None
    except Exception as e:
        print(f"Error: {e}")
        return None


if __name__ == "__main__":
    username = input("Username: ")
    password = input("Password: ")
    login(username, password)
