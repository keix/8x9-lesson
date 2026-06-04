import csv
import json

input_file = "input.csv"
output_file = "output.json"

items = []

# CSV を読み込む
with open(input_file, "r", encoding="utf-8") as f:
    reader = csv.DictReader(f)

    for row in reader:
        items.append(row)

# ロールの一覧を作る（重複を除いて、出てきた順を保つ）
roles = []
for item in items:
    if item["role"] not in roles:
        roles.append(item["role"])

# ロールを番号付きで表示する
print("ロールを選んでください:")
for i in range(len(roles)):
    print(str(i + 1) + ": " + roles[i])

# 番号を入力させる
number = int(input("番号: "))
selected_role = roles[number - 1]

# 選んだロールだけに絞る
filtered = []
for item in items:
    if item["role"] == selected_role:
        filtered.append(item)

# JSON に変換して書き出す
with open(output_file, "w", encoding="utf-8") as f:
    json.dump(filtered, f, ensure_ascii=False, indent=2)

print("変換しました:", output_file, "(role =", selected_role + ")")
