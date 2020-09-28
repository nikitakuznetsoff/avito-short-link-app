package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"shortlinkapp/pkg/database"
	"shortlinkapp/pkg/handlers"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gomodule/redigo/redis"
	"github.com/julienschmidt/httprouter"
)

func main() {
	dsn := "root:pass@tcp(db_mysql:3306)/balanceapp?"
	dsn += "&charset=utf8"
	dsn += "&interpolateParams=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	repo := database.NewRepo(db)
	templates := template.Must(template.ParseGlob("./templates/*"))
	handler := handlers.LinksHandler{Repo: repo, Tmpl: templates}

	router := httprouter.New()
	router.GET("/", handler.Index)
	router.GET("/:link", handler.GetLink)
	router.POST("/create", handler.CreateShortLink)

	address := ":9000"
	fmt.Printf("Starting server on port %s", address)
	log.Fatal(http.ListenAndServe(address, router))
}
