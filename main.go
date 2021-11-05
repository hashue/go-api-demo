package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "log"
    "fmt"
)

func getHello(w http.ResponseWriter, r *http.Request) {
   w.WriteHeader(http.StatusOK)
   fmt.Fprintf(w, "hello!") //Fprintf can specify writer as parameter.
}

func main() {
    var port string = ":8000"

    //initialize of router
    r := mux.NewRouter()

    //routing
    r.HandleFunc("/api/hello",getHello).Methods("GET")

    fmt.Printf("Listening on localhost%s...", port)
    
    log.Fatal(http.ListenAndServe(port,r))
}
