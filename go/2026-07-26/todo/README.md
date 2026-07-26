# TODO アプリ（Goal / Task CRUD）

Goal（目標）を決めて、それに紐づく Task を追加できる TODO アプリ。
[asteroid](../../2026-04-12/asteroid/) と同じ 3 層構造（Handler → Service → Store）で、**Service は全てスタブ**になっています。各メソッドの TODO コメントに沿って実装するのが課題です。

## 起動

```bash
go run ./cmd/server
```

起動時にサンプルデータ（`goal-1`, `task-1`）が Store に直接登録されます。

## エンドポイント

| Method | Path | 説明 |
| --- | --- | --- |
| POST | `/goals` | Goal 作成 |
| GET | `/goals` | Goal 一覧 |
| GET | `/goals/:id` | Goal 取得 |
| PUT | `/goals/:id` | Goal 更新 |
| DELETE | `/goals/:id` | Goal 削除（配下の Task もカスケード削除） |
| POST | `/goals/:id/tasks` | Goal に Task を追加 |
| GET | `/goals/:id/tasks` | Goal の Task 一覧 |
| GET | `/tasks/:id` | Task 取得 |
| PUT | `/tasks/:id` | Task 更新（title / done） |
| DELETE | `/tasks/:id` | Task 削除 |

## 動作確認用サンプル

```bash
# Goal を作成
curl -X POST http://localhost:8080/goals \
  -H 'Content-Type: application/json' \
  -d '{"title": "Learn Go", "description": "Build a layered web app"}'

# Goal 一覧
curl http://localhost:8080/goals

# Goal に Task を追加
curl -X POST http://localhost:8080/goals/goal-1/tasks \
  -H 'Content-Type: application/json' \
  -d '{"title": "Implement the goal service"}'

# Task を完了にする
curl -X PUT http://localhost:8080/tasks/task-1 \
  -H 'Content-Type: application/json' \
  -d '{"title": "Implement the goal service", "done": true}'

# Goal の Task 一覧
curl http://localhost:8080/goals/goal-1/tasks

# Goal を削除（Task もまとめて消える）
curl -X DELETE http://localhost:8080/goals/goal-1
```

※ Service がスタブのうちは、`/health` 以外は全て `{"error":"server_error"}` を返します。

## 構造

```
todo/
├── cmd/server/main.go        ← エントリポイント（配線 + サンプルデータ登録）
└── internal/
    ├── http/                 ← Handler 層（gin に依存してよい唯一の場所）
    │   ├── server.go             ルーティング + DI の配線
    │   ├── goal/handler.go       Goal エンドポイント
    │   └── task/handler.go       Task エンドポイント
    ├── todo/                 ← Service 層（ビジネスロジック / 純粋）★実装課題
    │   ├── goal/
    │   │   ├── service.go        Create / Get / List / Update / Delete（スタブ）
    │   │   ├── result.go         層をまたぐ返り値型
    │   │   └── errors.go         業務エラー種別
    │   └── task/
    │       ├── service.go        Create / Get / ListByGoal / Update / Delete（スタブ）
    │       ├── result.go
    │       └── errors.go
    └── store/                ← Store 層（永続化）
        ├── interfaces.go         契約（GoalStore / TaskStore interface）
        ├── entity/               ドメインのデータ型（Goal, Task）
        └── memory/               in-memory 実装（実装済み）
```

レイヤー構造の考え方は asteroid の [docs/structure.md](../../2026-04-12/asteroid/docs/structure.md) を参照。ポイントは同じ：

- 依存は一方通行（Handler → Service → Store）。Service は HTTP も map も知らない
- Service は `store.GoalStore` / `store.TaskStore` という **interface** にだけ依存（memory → SQL の差し替えが Service 無変更でできる）
- 返り値 `(*Result, ErrorType, error)` で **業務エラー**（invalid_request / not_found）と**インフラエラー**（500）を型で分離
- 配線は `internal/http/server.go` の `setupRoutes` に集約

## 実装課題

`internal/todo/goal/service.go` と `internal/todo/task/service.go` の TODO を上から順に埋めていく。

1. **Goal Create** — バリデーション → `uuid.NewString()` で ID 採番 → `SaveGoal`
2. **Goal Get / List** — `entity.ErrGoalNotFound` を `ErrorNotFound` に変換する（`errors.Is` を使う）
3. **Goal Update** — 既存を取得してから上書き。`CreatedAt` は維持、`UpdatedAt` だけ更新
4. **Goal Delete** — Goal 削除後に `DeleteTasksByGoal` でカスケード削除
5. **Task Create** — 親 Goal の存在確認が最初（存在しない Goal に Task をぶら下げない）
6. **Task ListByGoal / Get / Update / Delete** — 「空リスト」と「Goal が無い」は別の答え、に注意
