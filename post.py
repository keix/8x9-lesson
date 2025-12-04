import requests

text = input("Tweet: ")
res = requests.post("http://localhost:8080/tweets", json={"text": text})
print(res.text)
