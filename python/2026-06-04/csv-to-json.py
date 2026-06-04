# CSV を読み込んで JSON っぽい文字列として出力する
# ライブラリなし版

input_file = "input.csv"
output_file = "output.json"

# CSV を読む
with open(input_file, "r", encoding="utf-8") as f:
    lines = f.readlines()

# 先頭行をヘッダーとして使う
header_line = lines[0].strip()
headers = header_line.split(",")

# データ行を処理する
items = []

for line in lines[1:]:
    line = line.strip()

    # 空行は無視する
    if line == "":
        continue

    values = line.split(",")

    item = {}

    for i in range(len(headers)):
        key = headers[i]
        value = values[i]
        item[key] = value

    items.append(item)

# JSON 文字列を作る
json_text = "[\n"

for i in range(len(items)):
    item = items[i]

    json_text += "  {\n"

    keys = list(item.keys())

    for j in range(len(keys)):
        key = keys[j]
        value = item[key]

        json_text += '    "' + key + '": "' + value + '"'

        if j < len(keys) - 1:
            json_text += ","

        json_text += "\n"

    json_text += "  }"

    if i < len(items) - 1:
        json_text += ","

    json_text += "\n"

json_text += "]\n"

# JSON をファイルに書き出す
with open(output_file, "w", encoding="utf-8") as f:
    f.write(json_text)

print("変換しました:", output_file)
