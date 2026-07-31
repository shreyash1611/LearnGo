# Handlers- These handlers are designed to work as a functional listener, or a function

-Like if a Client Sends a Request
1. Go server listens on the port, handler receives client request
2. As per the server set up, it sends a response back. 
3. func name(w http.ResponseWriter, r *http.Request)- follows the syntax where "/" tells where the handler works when called(Home for example for localhost:8000/) 

# MUX- Fancy word for URL routers. When using http.HandleFunc, we pass what address the handler serves. 
1. For the same MUX, we can have multiple handlers explicitly to handle that address, just different pages. THE ROUTE ADDRESS ARE NOT CASE SENSITIVE    


# NOTE- Notice how "/" is a root path, where without handling it explicitly, all the pages display the same thing AKA subtree path. "/snippet/view" 


# RESTFUL APIS- We have something called read(r.Method) which is an enum of ("GET","POST","PUT","DELETE"). We can check the method that a handler expects, and handle error from here to ensure that we do not mess up the code

# AUTH CRITICAL 1: Changing the response header map after a call to w.WriteHeader() or w.Write() will have no effect on the headers that the user receives. You need to make sure that your response header map contains all the headers you want before you call these methods.