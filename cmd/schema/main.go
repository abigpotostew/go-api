package main

import (
	"github.com/abigpotostew/go-api/internal/database"
	"os"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")

	//pg := pg.Connect(&pg.Options{
	//	Network:  "tcp",
	//	Addr:     "",
	//	User:     "",
	//	Password: "",
	//	Database: "neondb",
	//	TLSConfig: &tls.Config{
	//		InsecureSkipVerify: true,
	//	},
	//})
	db, err := database.ConnectDatabase(dsn)
	if err != nil {
		panic(err)
	}
	err = database.CreateSchema(db)
	if err != nil {
		panic(err)
	}
}
