# Asteroid - OIDC Authorization Server (Learning Edition)

A learning implementation of an OIDC authorization server.

## Authorization Endpoint

### Request Parameters

| Parameter | Required | Description |
|-----------|----------|-------------|
| `client_id` | Yes | Client identifier |
| `redirect_uri` | Yes | Authorization code delivery URL (must exactly match registered URI) |
| `response_type` | Yes | Only `code` is supported |
| `scope` | Yes | Only `openid` is supported |
| `state` | Yes | Random string for CSRF protection |
| `nonce` | No | For ID token replay attack prevention |
| `code_challenge` | Required for public clients | PKCE challenge (Base64URL encoded) |
| `code_challenge_method` | Required for public clients | Only `S256` is supported |

### Headers

| Header | Required | Description |
|--------|----------|-------------|
| `X-Authenticated-User` | Yes | Authenticated user ID (typically set by reverse proxy) |

### Request Example

```bash
curl -X GET "http://localhost:8080/authorize?\
client_id=my-client&\
redirect_uri=http://localhost:3000/callback&\
response_type=code&\
scope=openid&\
state=abc123&\
nonce=xyz789" \
  -H "X-Authenticated-User: user123"
```

### Response

**Success:** HTTP 302 Redirect
```
Location: http://localhost:3000/callback?code=<authorization_code>&state=abc123
```

**Error:** HTTP 400 JSON
```json
{
  "error": "invalid_request"
}
```

### Error Codes

| Code | Description |
|------|-------------|
| `invalid_request` | Missing required parameter, missing state, or duplicate nonce |
| `invalid_client` | Client does not exist |
| `invalid_redirect_uri` | redirect_uri does not match registered URI |
| `unsupported_response_type` | response_type is not `code` |
| `invalid_scope` | scope is not `openid` |
| `access_denied` | User authentication failed |

## Directory Structure

```
internal/
├── store/                    # Data store layer
│   ├── interfaces.go         # Store interfaces
│   ├── entity/               # Entity definitions
│   │   ├── client.go         # Client
│   │   ├── authcode.go       # AuthCode
│   │   └── errors.go         # Error definitions
│   └── memory/               # In-memory implementation
│       ├── client.go
│       ├── authcode.go
│       └── nonce.go
├── userinfo/                 # User information provider
│   ├── provider.go           # Provider interface
│   └── source/
│       └── memory.go         # In-memory implementation
├── oidc/authorize/           # Authorization business logic
│   ├── errors.go             # ErrorType definition
│   ├── result.go             # Result struct
│   └── service.go            # Service (TODO implementation)
└── http/authorize/           # HTTP handler
    └── handler.go
```

## Learning Path

Implement the TODOs in `internal/oidc/authorize/service.go` in order:

1. **Step 1-2:** Validate required parameters
2. **Step 3:** Check nonce duplication
3. **Step 4-5:** Validate response_type / scope
4. **Step 6:** Validate client
5. **Step 7:** Validate PKCE
6. **Step 8:** Validate redirect_uri
7. **Step 9:** Verify user authentication
8. **Step 10-11:** Generate authorization code and redirect

## Running Tests

```bash
# See examples/authorize.sh
chmod +x examples/authorize.sh
./examples/authorize.sh
```

## Security Points

- **State:** Required to prevent CSRF attacks
- **Nonce:** Prevents ID token replay attacks
- **PKCE:** Prevents authorization code interception for public clients
- **Exact redirect_uri matching:** Prevents open redirector attacks
