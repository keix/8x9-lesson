import requests

text = input("Tweet: ")
res = requests.post("http://localhost:18080/tweets", json={"text": text})
print(res.text)
