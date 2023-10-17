package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    // Ambil nilai port dari ENV variabel APPPORT
    port := os.Getenv("APPPORT")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hallo dunia, Saya berjalan di port %s\n", port)
    })

    // Mulai server web pada port yang diambil dari ENV variabel
    fmt.Printf("Server berjalan di port %s\n", port)
    http.ListenAndServe(":" + port, nil)
}
