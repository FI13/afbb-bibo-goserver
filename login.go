package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/fi13/afbb-bibo-goserver/connection"
	"github.com/fi13/afbb-bibo-goserver/model"
	"github.com/fi13/afbb-bibo-goserver/util"
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
	ws.Route(ws.GET("/login").To(u.login))

	container.Add(ws)
}

func (u LoginResource) requestSalt(req *restful.Request, resp *restful.Response) {
	parameterMap, err := util.ExtractParameters(req)
	if err != nil {
		resp.WriteErrorString(http.StatusNotFound, err.Error())
		return
	}
	name := parameterMap["name"][0]
	log.Printf("request salt for user with name: %v", name)
	salt, err := connection.MySqlConnection.SelectStr("select Salt from benutzer where Name=?", name)
	if err == nil {
		resp.Write([]byte(salt))
		return
	}
	resp.WriteErrorString(http.StatusNotFound, err.Error())
}

func (u LoginResource) login(req *restful.Request, resp *restful.Response) {
	parameterMap, err := util.ExtractParameters(req)
	if err != nil {
		resp.WriteErrorString(http.StatusNotFound, err.Error())
		return
	}
	name := parameterMap["name"][0]
	hash := parameterMap["hash"][0]
	log.Printf("try to login user: %v", name)
	var curator model.Curator
	err = connection.MySqlConnection.SelectOne(&curator, "select * from benutzer where Name=?", name)
	if err == nil && curator.ValidateHash(hash) {
		log.Printf("login successfull")
		object, _ := json.Marshal(curator)
		resp.Write(append([]byte(randToken()+"\n"), object...))
		return
	}
	resp.WriteErrorString(http.StatusUnauthorized, err.Error())
}

func randToken() string {
	b := make([]byte, 40)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
