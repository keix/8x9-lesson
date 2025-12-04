from http.server import ThreadingHTTPServer, BaseHTTPRequestHandler
import json

TWEETS = []

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

    def do_GET(self):
        if self.path == "/tweets":
            return self._send_json(TWEETS)
        self._send_json({"error": "not found"}, 404)

    def do_POST(self):
        if self.path != "/tweets":
            return self._send_json({"error": "not found"}, 404)

        length = int(self.headers.get("Content-Length", 0))
        body = self.rfile.read(length)
        data = json.loads(body)

        tweet = {"id": len(TWEETS) + 1, "text": data["text"]}
        TWEETS.append(tweet)

        return self._send_json(tweet, 201)


def run():
    print("Server running at http://localhost:8080")
    server = ThreadingHTTPServer(("0.0.0.0", 8080), Handler)
    server.serve_forever()

if __name__ == "__main__":
    run()


