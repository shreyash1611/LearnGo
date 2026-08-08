# IMPORTANT — ListenAndServe blocks
- `http.ListenAndServe` starts the server and **does not return** under normal use.
- Any `HandleFunc` / setup lines **below** it in `main` never run.
- Always: create mux → register routes → then `ListenAndServe`.
- If you see Go's default `404 page not found` for every URL, you probably registered routes after listen (empty mux).

# Handlers- These handlers are designed to work as a functional listener, or a function

-Like if a Client Sends a Request
1. Go server listens on the port, handler receives client request
2. As per the server set up, it sends a response back. 
3. func name(w http.ResponseWriter, r *http.Request)- follows the syntax where "/" tells where the handler works when called(Home for example for localhost:8000/) 

# MUX- Fancy word for URL routers. When using http.HandleFunc, we pass what address the handler serves. 
1. For the same MUX, we can have multiple handlers explicitly to handle that address, just different pages. THE ROUTE ADDRESS ARE CASE SENSITIVE    


# NOTE- Notice how "/" is a root path, where without handling it explicitly, all the pages display the same thing AKA subtree path. "/snippet/view" is a specific static path


# RESTFUL APIS- We have something called read(r.Method) which is an enum of ("GET","POST","PUT","DELETE"). We can check the method that a handler expects, and handle error from here to ensure that we do not mess up the code

# AUTHOR CRITICAL 1: Changing the response header map after a call to w.WriteHeader() or w.Write() will have no effect on the headers that the user receives. You need to make sure that your response header map contains all the headers you want before you call these methods.

# AUTHOR CRITICAL 2: We can change certain constants like "POST" a string to "http.MethodPost" and still achieve the same thing. THis makes it less prone to errors, and better documentibility.


# RESPONSE DATA- ADDS CONTENT TYPE, DATE, CONTENT LENGTH
<!-- shreyashchaurasia@shreyashchaurasia-HP-Laptop-14s-dr2xxx:~/Documents/learngo/backend$ curl -i -X "POST"  http://localhost:4000/snippet/create -->
<!-- HTTP/1.1 200 OK -->
<!-- Date: Sun, 02 Aug 2026 09:15:08 GMT -->
<!-- Content-Length: 18 -->
<!-- Content-Type: text/plain; charset=utf-8 #DEFAULT-application/octet-stream -->
<!--  -->
<!-- Creating a snippet -->

# ^ NOTE FOR THE TOPIC ABOVE You should set Content-Type yourself when you care about the type (JSON, HTML forms, etc.). Sniffing is a fallback for demos/plain text.-
<!-- w.Header().Set("Content-Type", "text/plain; charset=utf-8") -->
<!-- w.Write([]byte("Creating a snippet")) -->
<!-- Order still matters: set Content-Type
 before Write / WriteHeader. -->


# THE RESPONSE- 
1. STATUS- 
2. HEADERS- CONTENTS TYPES, ALLOW, LENGTH, DATE, CACHE-CONTROL ETC ARE HEADERS TO DECIDE IMP INFO
3. BODY- RESPONSE WRITES ARE DONE HERE

2a- Headers are map of string vectors or rather header := map string[]. add adds as a new value to the key, whereas set makes it all in 1 string and adds as is. 
2b.HTTP treats header names as case-insensitive: content-type, Content-Type, CONTENT-TYPE-Those are the same header. But a Go map is case-sensitive — "content-type" and "Content-Type" would be two different keys if nothing cleaned them up.Canocalization basically standardises it.


# STRUCTURING THE FILES
- We split the app top-down: `cmd/` = runnable programs, `internal/` = private project packages.
- `internal` is special in Go: packages under it can ONLY be imported by code inside the parent of that `internal` folder (for us: inside `backend/` / our module). Outside projects cannot import them.
- Same-folder files in `package main` (e.g. `main.go` + `handlers.go`) see each other automatically — run with `go run .`, not `go run main.go`.