package main

import (
    "fmt"
    "net/http"
    "time"
)

var frames = []string{
    "\033[31mhmm\033[0m   ",
    " \033[32mhmm\033[0m  ",
    "  \033[33mhmm\033[0m ",
    "   \033[34mhmm\033[0m",
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    flusher, ok := w.(http.Flusher)
    if !ok {
        http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
        return
    }

    // Infinite loop to keep animation running until client closes connection
    for i := 0; ; i++ {
        fmt.Fprintf(w, "\033[2J\033[H%s\n", frames[i%len(frames)])
        flusher.Flush()
        time.Sleep(200 * time.Millisecond)
    }
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Colorful parrot running at http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Server error:", err)
    }
}
