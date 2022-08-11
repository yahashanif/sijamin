package main

import (
	"rest-api-sijamin/db"
	"rest-api-sijamin/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Start(":9000")
}
