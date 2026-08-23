# HTTP explained

Longer explanations for topics in [`Notes.md`](Notes.md).  
**Main go-to (story + functions):** open [`SignatureIdea.html`](SignatureIdea.html) in a browser.

| File | Role |
|------|------|
| SignatureIdea.html | Main guide — read this first |
| Notes.md | Your short ideas |
| HTTP-EXPLAINED.md | This file — deeper “why” |

---

## ListenAndServe blocks (see Notes: IMPORTANT)

`ListenAndServe` binds the port and loops forever. Lines **below** it in `main` do not run.

Wrong: listen first, then `HandleFunc` → empty mux → Go’s default `404 page not found`.  
Right: mux → register all routes → `ListenAndServe` last.

Your custom `notfound` says something like “404 Page Not LOL Found”. Default wording means the handler never ran.

`go ListenAndServe` then register = race. Don’t. No need for goroutines here; the book doesn’t teach that in this chapter.

---

## Handlers (see Notes: Handlers)

```go
func home(w http.ResponseWriter, r *http.Request)
```

| Piece | Role |
|-------|------|
| `w` | Write status, headers, body |
| `r` | Read method, path, query, body |
| return | None — reply is written to `w` |

Flow: client → server → mux picks handler → handler writes response.

---

## Mux / router (see Notes: MUX)

Mux only chooses **which** handler runs.

- Paths are case-sensitive
- Bare `"/"` can be a catch-all; `/{$}` is exact `/` only (Go 1.22+)
- Unknown path → **404**; wrong method on a known path → **405** + `Allow`

---

## Query string (see Notes: Query string)

URL parts:

```text
http://localhost:4000/snippet/view?id=12
       └──── host ────┘└─── path ────┘└query┘
```

Mux matches **path**. Query is read in the handler:

```go
r.URL.Query().Get("id")  // string; "" if missing / empty
strconv.Atoi(...)        // string → int + error
```

Must use `id=12` (`=`). `id-12` is not “id equals 12”.  
`?id`, `?id=`, missing → empty → treat as 404 for view.

---

## Methods (see Notes: RESTFUL)

Prefer `http.MethodGet` / `http.MethodPost`.  
Wrong method → headers first (`Allow`), then `WriteHeader(405)`, then body.

---

## Response = status + headers + body (see Notes: THE RESPONSE)

```
STATUS   → 200 / 404 / 405 …
HEADERS  → Content-Type, Allow, Date…
BODY     → text / JSON / HTML from template
```

**Author Critical 1:** set headers **before** `WriteHeader` or `Write`.

`Set` vs `Add`: map is always `map[string][]string`. `Set` replaces; `Add` appends.  
`Set`/`Get`/… canonicalize names (`content-type` → `Content-Type`).

---

## Auto headers & Content-Type (see Notes: RESPONSE DATA)

Go may add `Date`, `Content-Length`, `Content-Type` (sniff). Prefer setting `Content-Type` yourself.  
`http.Error` forces `text/plain`.

---

## Templates — CRITICAL (see Notes: Templates)

**Main go-to for this topic:** [`SignatureIdea.html`](SignatureIdea.html) → Deep dive: HTML templates.

### Idea

Don’t hardcode full HTML in Go. Store `.tmpl` files; Go parses them and writes the result into `w` (same as body via `Write`, but assembled from files).

### `define` vs `template` (lowest level)

| Action | Job | Analogy |
|--------|-----|---------|
| `{{define "title"}}Home{{end}}` | **Save** clip named `"title"` | `title = "Home"` |
| `{{template "title" .}}` | **Paste** that clip here | `print(title)` |

`define` alone shows nothing. `template` is what inserts into the output.

### File roles

| File | Role |
|------|------|
| `base.tmpl` | Shared layout; `define "base"`; pastes title / nav / main |
| `partials/nav.tmpl` | Shared nav; `define "nav"` |
| `pages/home.tmpl` | Home only; `define "title"` + `define "main"` |
| later `pages/view.tmpl` | View only; its own `"title"` + `"main"` |

### CRITICAL — one page per ParseFiles

After `ParseFiles`, all `define` names live in **one bag**.

```text
home()  → ParseFiles(base, nav, home.tmpl)     // OK
view()  → ParseFiles(base, nav, view.tmpl)     // OK

BAD: home() ParseFiles(..., home.tmpl, view.tmpl)
     both define "title" → last file wins (silent overwrite)
```

Not loading `view.tmpl` inside `home()` is correct.

| Situation | Result |
|-----------|--------|
| `base` pastes `"title"` but no `define "title"` in the set | `ExecuteTemplate` error |
| Wrong/missing file path | `ParseFiles` error |
| Two `"title"` defines in same ParseFiles | Last wins — wrong page title possible |
| `view.tmpl` on disk but not in `home`’s files | Fine |

### Handler wiring (no magic)

```go
// main.go — URL → Go function
homemux.HandleFunc("/{$}", home)

// home() — which files to load; start render at "base"
tmpl, err := template.ParseFiles(base, homePage, nav)
err = tmpl.ExecuteTemplate(w, "base", nil)  // error only; nil data for now
```

Go does not choose `home.tmpl` because of the filename. **You** listed it inside `home`.

### cwd / links / methods

- Paths relative to **cwd** → run from `backend/`: `go run ./cmd/web`
- `href="/"` → same host:port as the address bar (no `:4000` in HTML)
- Typing a URL is often GET — **not always**. Forms/`curl` can POST etc. Check `r.Method` when needed.

Don’t `w.Write` extra text after a successful `ExecuteTemplate` unless you want junk after `</html>`.

---

## Project layout & `internal` (see Notes: STRUCTURING)

```text
backend/                 ← go.mod (module learngo)
  cmd/web/               ← package main
  ui/html/pages/         ← templates
  internal/              ← private packages (later)
```

`internal` = special **directory name**, not a language keyword. Only code under its parent (`backend/`) may import it.

`go run main.go` alone misses `handlers.go`. Use `go run ./cmd/web` from `backend/` or `go run .` from `cmd/web`.

---

## Handy commands

```bash
cd backend && go run ./cmd/web
curl -i http://localhost:4000/
curl -i 'http://localhost:4000/snippet/view?id=12'
curl -i -X POST http://localhost:4000/snippet/create
```
