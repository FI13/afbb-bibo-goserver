package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/fi13/afbb-bibo-goserver/connection"
	"github.com/fi13/afbb-bibo-goserver/model"
)

type UserResource struct { /* empty */
}

func (r UserResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/user").
		Doc("manages user").
		Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	ws.Route(ws.GET("/getBorrowers").To(r.getBorrowers).Param(ws.HeaderParameter("sessionId", "token").
		DataType("string")))
	// ws.Route(ws.GET("/login").To(r.login))

	container.Add(ws)
}

func (u UserResource) getBorrowers(req *restful.Request, resp *restful.Response) {
	if !ValidateSession(req, resp) {
		return
	}
	var borrowers []model.Borrower
	_, err := connection.MySqlConnection.Select(&borrowers, "select * from ausleiher")
	if err != nil {
		resp.WriteErrorString(http.StatusNotFound, err.Error())
		log.Printf(err.Error())
		return
	}

	for i := range borrowers {
		object, _ := json.Marshal(&borrowers[i])
		resp.Write(object)
		resp.Write([]byte("\n"))
	}
	return
}
