package main

import (
	"crypto/rsa"
	"database/sql"

	"log"
	"net/http"
	"os"
	"time"

	"bitbucket.org/sunchero/auth/handler"
	"bitbucket.org/sunchero/auth/service"

	_ "github.com/lib/pq"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
	db        *sql.DB
)

func main() {
	//init db , rsa/jwt , goth
	boot()
	// create backend
	db, err := sql.Open("postgres", os.Getenv("DBURL"))
	if err != nil {
		log.Fatalf(`could not init db %v`, err)
	}
	s := service.New(db, signKey)
	//create http handler
	router := handler.New(s)
	// init Servcer
	server := http.Server{
		Addr:              "0.0.0.0:3000",
		Handler:           router,
		ReadHeaderTimeout: time.Second * 5,
		ReadTimeout:       time.Second * 15,
	}
	log.Fatal(server.ListenAndServe())
	defer db.Close()
}
