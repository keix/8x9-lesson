import requests

res = requests.get("http://localhost:18080/tweets")
for t in res.json():
    print(f"{t['id']}. {t['text']}")
