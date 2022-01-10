package main

import (
	"yawoenapi/src/database"
	"yawoenapi/src/readCSV"
	"yawoenapi/src/routes"
)

func main() {
	database.InitDatabase()
	readCSV.ReadCSVInsert("./media/q1_catalog.csv")

	routes.InitRoute()
}
