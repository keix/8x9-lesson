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

# JSON に変換して書き出す
with open(output_file, "w", encoding="utf-8") as f:
    json.dump(items, f, ensure_ascii=False, indent=2)

print("変換しました:", output_file)
