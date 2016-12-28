package main

import (

    "fmt"
    "net/http"
    "encoding/json"
    //"html   log"
    //Imported from Go/src/ 
    //Current directory:"./models"
)

func main() {

   fmt.Println("Server Running on localhost:8080")

   http.HandleFunc("/",defaultHandler)
   http.HandleFunc("/about/",aboutHandler)
   http.ListenAndServe(":8080",nil)
}

func defaultHandler(w http.ResponseWriter, r *http.Request){
   fmt.Fprintf(w,"Hello, %s",r.URL.Path[1:])
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
    m := Message{"Welcome to about page"}
    b, err := json.Marshal(m)

    if err!= nil{
        panic(err)
    }

    w.Write(b)
}

type Message struct{
    Text string
}