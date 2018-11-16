package main

import (
    _ "go_api/database"
    "go_api/router"
    db "go_api/database"
)

func main() {
    defer db.PostgresDb.Close()
    router := router.InitRouter()
    router.Run(":8000")
}
