package main

import (
	"yawoenapi/src/controllers"
	"yawoenapi/src/database"
)

func main() {
	database.InitDatabase()
	controllers.ReadCSVInsert("./media/q1_catalog.csv")
	controllers.InitRoute()
}
