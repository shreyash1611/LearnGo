# How these files work
- **SignatureIdea.html** = MAIN go-to (open in browser). Story of a request + what each call does.
- **Notes.md** (this file) = my short ideas / flashcards.
- **HTTP-EXPLAINED.md** = longer explanations of the same topics.

---

# IMPORTANT — ListenAndServe blocks
- `http.ListenAndServe` starts the server and **does not return** under normal use.
- Any `HandleFunc` / setup lines **below** it in `main` never run.
- Always: create mux → register routes → then `ListenAndServe`.
- Default `404 page not found` on every URL = routes registered after listen (empty mux).
- My custom `notfound` text is different — if I see the default wording, handler never ran.
- Do NOT use `go ListenAndServe` then register routes (race). Book doesn’t need `go` here.

# Handlers
- Handlers = functions that listen for a matched request and write a response.
- Client sends request → server on a port → handler runs → response back.
- `func name(w http.ResponseWriter, r *http.Request)` — no return value; answer goes into `w`.
- Path like `/` is registered on the mux; that’s when `home` runs (e.g. localhost:4000/).

# MUX
- Fancy word for URL router. `HandleFunc(path, handler)` = this address → that function.
- Same mux, many handlers, different paths. Paths are CASE SENSITIVE.
- Mux only picks the handler. It does not build the body.

# NOTE — `/` vs exact paths
- Bare `"/"` is a subtree / catch-all if you use it that way.
- `/{$}` = exact home only (Go 1.22+). Better for real 404s.
- `/snippet/view` = specific static path.

# Query string (`?id=12`)
- Mux looks at PATH only. `?id=12` is NOT part of the route.
- Need `=` not `-` → `?id=12` works, `?id-12` does not.
- `Get("id")` → string. Missing / `?id` / `?id=` → `""` → Atoi fails → 404.
- Valid view needs a real positive id (book: `err != nil || id < 1`).

# RESTFUL / methods
- `r.Method` is GET, POST, PUT, DELETE…
- Prefer `http.MethodGet` / `http.MethodPost` over raw strings.
- Wrong method → **405** + `Allow` header, not 404.
- Unknown path → **404**.

# AUTHOR CRITICAL 1
- After `WriteHeader` or `Write`, changing headers has no effect for the client.
- Set all headers you want BEFORE those calls.

# AUTHOR CRITICAL 2
- `"POST"` → `http.MethodPost` (same idea, fewer typos, clearer).

# RESPONSE DATA — auto headers
- Go may add Date, Content-Length, Content-Type.
- Fallback sniff; if unsure → `application/octet-stream`.
- Set Content-Type yourself for JSON/HTML when you care.
- `http.Error` forces `text/plain` and can wipe your JSON Content-Type.
- Order: Set Content-Type → WriteHeader → Write.

# THE RESPONSE
1. STATUS — 200, 404, 405…
2. HEADERS — Content-Type, Allow, Length, Date, Cache-Control…
3. BODY — what `Write` / `Fprintf` / template `Execute` send

2a — Header map is always `map[string][]string`. `Set` = replace with one value; `Add` = append to the list.
2b — Header names case-insensitive via Set/Add/Get (canonicalization). Direct map assign skips that.

# Templates (CRITICAL)
- HTML on disk → Go ParseFiles → ExecuteTemplate into `w`. Not giant HTML strings in Go.
- **`define`** = SAVE a clip under a name (does not show anything alone).
- **`template "name"`** = PASTE that named clip **here**.
- `base.tmpl` = shared layout (holes for title/nav/main). Page files like `home.tmpl` / later `view.tmpl` = only that page’s plugs.
- `nav.tmpl` = shared partial; also just a `define "nav"`.
- **CRITICAL:** each handler ParseFiles only `base` + `nav` + **that page’s** tmpl. Never dump every page into one ParseFiles.
- If `home.tmpl` and `view.tmpl` both `define "title"` in the **same** ParseFiles → **last one wins** (silent overwrite). Wrong title, no loud “duplicate” error.
- Not parsing `view.tmpl` inside `home()` is **correct** — home doesn’t need it.
- Missing `define` that `base` tries to `{{template}}` → **ExecuteTemplate error**.
- Missing/wrong file path → **ParseFiles error**.
- Go does not pick `home.tmpl` by magic: `main` maps `/` → `home()`; `home()` lists the files.
- `ExecuteTemplate(w, "base", nil)` — start at `"base"`; `nil` = no data yet; returns only `error`.
- cwd-relative paths; run from `backend/`. `href="/"` → same host:port (no :4000 in HTML).
- Requests are **not** always GET — check `r.Method` when it matters.
- Deep dive: SignatureIdea.html → “Deep dive: HTML templates”.

# STRUCTURING THE FILES
- Top-down: `cmd/` = runnable apps, `internal/` = private packages.
- `internal` = special folder name (not a language keyword). Only code under parent of `internal` (our `backend/`) can import it. Outsiders can’t.
- Same `package main` folder: `main.go` + `handlers.go` — run `go run .` or `go run ./cmd/web`, not only `go run main.go`.
