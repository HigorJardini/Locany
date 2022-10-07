package main

import (
    "internal/app"
)

func main() {
    // Start Fiber API server.
    app.sender()
    log.Fatal(app.Listen(":3000"))
}