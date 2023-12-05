This program creates a simple web server with form processing and static file
serving capabilities. The server responds to requests on `/form` by processing
form data and printing it, and on `/hello` by responding with "hello".

Additonally, static files can be served from the "./static" directory.

### 1. Imports
```go
import (
  "fmt"
  "log"
  "net/http"
)
```
The program imports necassary packages, `fmt` for formatted I/O,
`log` for logging and `net/http` for HTTP server functionality.

### 2. Form Handler
```go
func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
    	fmt.Fprintf(w, "ParseForm() err: %v", err)
    	return
    }
    fmt.Fprint(w, "POST request successful")
    name := r.FormValue("name")
    address := r.FormValue("address")
    fmt.Fprintf(w, "Name = %s\n", name)
    fmt.Fprintf(w, "Address = %s\n", address)
}
```
The `formHandler` function is responsible for handling the `/form` endpoint.
It parses the form data from the request and prints the values of the "name" 
and "address" fields.

### 3. Hello Handler
```go
func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
      	http.Error(w, "404 not found", http.StatusNotFound)
      	return
    }
    if r.Method != "GET" {
      	http.Error(w, "method is not supported", http.StatusNotFound)
      	return
    }
    fmt.Fprintf(w, "hello")
}
```
The `helloHandler` function handles the `/hello` endpoint. It responds with
"hello" only if the HTTP method is a GET request, otherwise, it returns a 
method is not allowed error.

### 4. Main Function
```go
func main() {
  	// Serving static files from the "./static" directory
  	fileServer := http.FileServer(http.Dir("./static"))
  	http.Handle("/", fileServer)

  	// Registering the formHandler and helloHandler for specific endpoints
  	http.HandleFunc("/form", formHandler)
  	http.HandleFunc("/hello", helloHandler)

  	// Starting the server on port 8080
  	fmt.Printf("Starting server at port 8080\n")
  	if err := http.ListenAndServe(":8080", nil); err != nil {
    		log.Fatal(err)
  	}
}
```
In the `main` function:
- A file server is setup to serve the static files from the "./static" directory.
- Handlers for `/form` and `/hello` endpoints are registered using `htpp.HandleFunc`.
- The server is started on port 8080 using `http.ListenAndServe`, and any errors are logged.
