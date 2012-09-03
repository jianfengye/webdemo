package main

import (
    "net/http"
    "log"
)

func main() {
    log.Println("main")
    http.Handle("/css/", http.FileServer(http.Dir("template")))
    http.Handle("/js/", http.FileServer(http.Dir("template")))
    
    http.HandleFunc("/admin/", adminHandler)
    http.HandleFunc("/login/",loginHandler)
    http.HandleFunc("/ajax/",ajaxHandler)
    http.HandleFunc("/",NotFoundHandler)
    http.ListenAndServe(":8888", nil)
}