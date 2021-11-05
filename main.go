package main

import (
    "net/http"
    "fmt"
    "os"
    "os/signal"
)

func getHello(w http.ResponseWriter, r *http.Request) {

   // Varidate Method
   if r.Method != http.MethodGet{
       w.WriteHeader(http.StatusMethodNotAllowed)
       w.Write([]byte("This api support only GET method"))
       return
   }

   w.WriteHeader(http.StatusOK)
   fmt.Fprintf(w, "hello!") //Fprintf can specify writer as parameter.
}

func siginthandler(){

    fmt.Println("\nStop server...")

}

func main() {
    var port string = ":8000"

    //routing
    http.HandleFunc("/api/hello",getHello)

    fmt.Printf("Listening on localhost%s...", port)
    
    go http.ListenAndServe(port,nil)

    quit := make(chan os.Signal)

    signal.Notify(quit, os.Interrupt)

    <-quit

    siginthandler();

}
