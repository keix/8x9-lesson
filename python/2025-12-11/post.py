import requests

token = input("Token: ")
text = input("Tweet: ")

headers = {"Authorization": f"Bearer {token}"}
res = requests.post(
    "http://localhost:18080/tweets", json={"text": text}, headers=headers
)
print(res.text)
