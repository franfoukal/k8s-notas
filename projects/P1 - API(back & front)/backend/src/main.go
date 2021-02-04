package main

import (
    // "fmt"
    "net/http"
    "time"
    "os"
    "encoding/json"
)

type HandsOn struct {
    Time time.Time `json:"time"`
    Hostname string `json:"hostname"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
    
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    resp := HandsOn{
        Time: time.Now(),
        Hostname: os.Getenv("HOSTNAME"),
    }
    jsonResponse, err := json.Marshal(&resp)
    if err != nil {
        w.Write([]byte("Error"))
        return
    }
    // w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write(jsonResponse)
}

func main() {
    http.HandleFunc("/", ServeHTTP)
    http.ListenAndServe(":9090", nil)
}