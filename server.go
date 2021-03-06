package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/emicklei/go-restful"
	"github.com/fi13/afbb-bibo-goserver/connection"
)

import _ "github.com/go-sql-driver/mysql"

func main() {
	conn := connection.Setup()
	defer conn.Db.Close()

	wsContainer := restful.NewContainer()
	login := LoginResource{}
	login.Register(wsContainer)
	user := UserResource{}
	user.Register(wsContainer)

	port := 8080
	log.Printf("start listening on localhost:%v", port)
	server := &http.Server{Addr: ":" + strconv.Itoa(port), Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
