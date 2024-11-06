package main

import (
    "fmt"
    "net/http"
    "os"
)

type App struct {
    name string
    config *Config
}

type Config struct {
    Port string
}

func (app *App) handler(w http.ResponseWriter, r *http.Request) {
    // This will cause a panic if config is nil
    fmt.Fprintf(w, "App Name: %s, Port: %s", app.name, app.config.Port)
}

func main() {
    app := &App{
        name: "MyApp",
        // config is intentionally not initialized to simulate a nil pointer dereference
        config: nil,
    }

    http.HandleFunc("/", app.handler)
    fmt.Println("Starting server on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Fprintln(os.Stderr, "Server error:", err)
    }
}
