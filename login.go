package main

import (
	"github.com/emicklei/go-restful"
	"github.com/fi13/afbb-bibo-goserver/connection"
	"github.com/fi13/afbb-bibo-goserver/util"
	"log"
	"net/http"
)

type LoginResource struct { /* empty */
}

func (u LoginResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/login").
		Doc("manages user session").
		Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	ws.Route(ws.GET("/requestSalt").To(u.requestSalt))

	container.Add(ws)
}

func (u LoginResource) requestSalt(req *restful.Request, resp *restful.Response) {
	nameMap, err := util.ExtractParameters(req)
	if err != nil {
		resp.WriteErrorString(http.StatusNotFound, err.Error())
		return
	}
	name := nameMap["name"][0]
	log.Printf("request salt for user with name: %v", name)
	hash, err := connection.MySqlConnection.SelectStr("select Salt from benutzer where Name=?", name)
	util.ServiceResponse(resp, hash, err)
}
