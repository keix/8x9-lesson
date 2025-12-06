# Server Setup

## Installation
Install with the following:
```bash
pip install requests
```

## Running the Server
```bash
python server.py
```

The server will start on port 18080.

## Login
First, login to get a token:
```bash
curl -X POST http://localhost:18080/login \
  -H "Content-Type: application/json" \
  -d '{"username": "user1", "password": "pass1"}'
```

Default users: `user1/pass1` and `admin/123`

## Posting
Post tweets using the token:
```bash
curl -X POST http://localhost:18080/tweets \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{"text": "Hello World!"}'
```

## View Tweets
```bash
curl http://localhost:18080/tweets
```

With just this setup, you can see the structure of a simple Twitter-like application with authentication.