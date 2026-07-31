# Handlers- These handlers are designed to work as a functional listener, or a function

## Like if a Client Sends a Request
1. Go server listens on the port, handler receives client request
2. As per the server set up, it sends a response back. 
3. func name(w http.ResponseWriter, r *http.Request)- follows the syntax where "/" tells where the handler works when called(Home for example for localhost:8000/) 