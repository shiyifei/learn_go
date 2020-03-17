package main

import (
 	db "use_gin/app/database"
	"use_gin/app"
)

func main() {
	defer db.SqlDB.Close()
	router := app.InitRouter()
	router.Run(":8100")
}
