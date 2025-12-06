from http.server import ThreadingHTTPServer, BaseHTTPRequestHandler
import json

TWEETS = []
USERS = {"admin": "password"}  # ID / Password は好きなのに変えてください
SESSIONS = {}


class Handler(BaseHTTPRequestHandler):
    def _send_json(self, obj, status=200):
        data = json.dumps(obj).encode()
        self.send_response(status)
        self.send_header("Content-Type", "application/json")
        self.send_header("Content-Length", str(len(data)))
        self.send_header("Connection", "close")
        self.end_headers()
        self.wfile.write(data)
        self.close_connection = True

    def _get_session_user(self):
        auth = self.headers.get("Authorization")
        if auth and auth.startswith("Bearer "):
            session_id = auth[7:]
            return SESSIONS.get(session_id)
        return None

    def do_GET(self):
        if self.path == "/tweets":
            return self._send_json(TWEETS)
        self._send_json({"error": "not found"}, 404)

    def do_POST(self):
        length = int(self.headers.get("Content-Length", 0))
        body = self.rfile.read(length)
        data = json.loads(body)

        if self.path == "/login":
            username = data.get("username")
            password = data.get("password")

            if username in USERS and USERS[username] == password:
                session_id = f"session_{len(SESSIONS) + 1}"
                SESSIONS[session_id] = username
                return self._send_json({"token": session_id, "user": username})
            else:
                return self._send_json({"error": "invalid credentials"}, 401)

        elif self.path == "/tweets":
            user = self._get_session_user()
            if not user:
                return self._send_json({"error": "login required"}, 401)

            tweet = {"id": len(TWEETS) + 1, "user": user, "text": data["text"]}
            TWEETS.append(tweet)
            return self._send_json(tweet, 201)

        return self._send_json({"error": "not found"}, 404)


def run():
    print("Server running at http://localhost:18080")
    server = ThreadingHTTPServer(("127.0.0.1", 18080), Handler)
    server.serve_forever()


if __name__ == "__main__":
    run()
