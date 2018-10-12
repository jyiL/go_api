package main

import (
    db "go_api/database"
)

func main() {
    defer db.PostgresDb.Close()
    // defer db.Mongo.Close()
    router := initRouter()
    router.Run(":8000")
}
