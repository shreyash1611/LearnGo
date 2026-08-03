# HTTP explained (companion to Notes.md)

Your own short notes live in [`Notes.md`](Notes.md). This file expands those same topics with more explanation.

---

## Handlers (see Notes: Handlers)

A handler is a function Go calls when a request matches a route.

```go
func home(w http.ResponseWriter, r *http.Request)
```

| Piece | Role |
|-------|------|
| `w` | Response writer — you send status, headers, body through this |
| `r` | Incoming request — method, path, headers, body |
| return type | None. The “answer” is written to `w`, not returned |

Flow: client request → server → mux picks handler → handler writes response.

---

## Mux / router (see Notes: MUX)

Mux = URL router. It only chooses **which handler** runs. It does not write the body.

- Register: `mux.HandleFunc("/snippet/view", snippetView)`
- Paths are case-sensitive
- Bare `"/"` is a **subtree / catch-all** (as in your Notes)
- `/{$}` (Go 1.22+) matches **only** exact `/` — better when you want real 404s for unknown paths

| Situation | Status |
|-----------|--------|
| No route matches | **404** (mux, if `/` is not a catch-all) |
| Route exists, wrong method | **405** + usually `Allow` header |

---

## Methods (see Notes: RESTFUL APIS)

`r.Method` is `"GET"`, `"POST"`, etc. Prefer `http.MethodGet`, `http.MethodPost` (your Author Critical 2).

Wrong method → set `Allow`, then `WriteHeader(405)`, then write body. Prefer that over sending **404**.

---

## Response = status + headers + body (see Notes: THE RESPONSE)

```
STATUS   → 200 OK / 404 / 405 …
HEADERS  → labels about the body (Content-Type, Allow, Date…)
BODY     → the actual content (text, JSON, HTML)
```

**Author Critical 1:** set all headers **before** `WriteHeader` or `Write`. After either runs, header changes are ignored by the client.

Order: `Header().Set(...)` → `WriteHeader(code)` → `Write(body)`.

---

## Auto headers & Content-Type (see Notes: RESPONSE DATA)

Go often adds `Date`, `Content-Length`, and `Content-Type` for you.

If you don’t set `Content-Type`, Go sniffs the body (`DetectContentType`). If unsure → `application/octet-stream`. Your curl showed `text/plain` for a normal sentence — expected.

Set it yourself when it matters:

```go
w.Header().Set("Content-Type", "application/json; charset=utf-8")
```

Note: `http.Error` forces `text/plain` and can overwrite a JSON `Content-Type` you set earlier.

---

## Header map (see Notes: 2a / 2b)

Type is always:

```go
map[string][]string
```

| Method | Effect |
|--------|--------|
| `Set` | Replace that key’s list with **one** value |
| `Add` | Append another value to the list |
| `Del` | Remove the key |
| `Get` | First value as `string` |
| `Values` | Full `[]string` |

**Canonicalization:** `Set`/`Add`/`Get`/… normalize names (`content-type` → `Content-Type`), so those APIs are case-insensitive. Direct assign `w.Header()["X-Foo"] = ...` skips that (rarely needed).

Request headers → `r.Header`. Response headers → `w.Header()`.

---

## Handy curls

```bash
cd backend && go run ./cmd/server
curl -i http://localhost:4000/
curl -i http://localhost:4000/snippet/view
curl -i -X POST http://localhost:4000/snippet/create
```
