package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/fi13/afbb-bibo-goserver/connection"
	"github.com/fi13/afbb-bibo-goserver/model"
	"github.com/fi13/afbb-bibo-goserver/util"
)

type LoginResource struct { /* empty */
}

type Session struct {
	Id   int
	Name string
	Time time.Time
}

var sessionMap = make(map[string]Session)

func (r LoginResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/login").
		Doc("manages user session").
		Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	ws.Route(ws.GET("/requestSalt").To(r.requestSalt))
	ws.Route(ws.GET("/login").To(r.login))

	container.Add(ws)
}

func ValidateSession(req *restful.Request, resp *restful.Response) bool {
	log.Printf("validate session, reading token...")
	parameterMap, err := util.ExtractParameters(req)
	if err != nil {
		resp.WriteErrorString(http.StatusUnauthorized, err.Error())
		return false
	}
	//util.LogAsJson(parameterMap)
	//sessionId := parameterMap["sessionId"][0]
	sessionId := parameterMap.Get("sessionId")
	log.Printf("validate session token: %v", sessionId)
	session := sessionMap[sessionId]
	// TODO validate time
	if &session == nil {
		resp.WriteErrorString(http.StatusUnauthorized, err.Error())
		log.Printf("unknown session token")
		return false
	}
	log.Printf("session OK")
	return true
}

func (u LoginResource) requestSalt(req *restful.Request, resp *restful.Response) {
	parameterMap, err := util.ExtractParameters(req)
	if err != nil {
		resp.WriteErrorString(http.StatusNotFound, err.Error())
		return
	}
	name := parameterMap.Get("name")
	//name := parameterMap["name"][0]
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
	name := parameterMap.Get("name")
	hash := parameterMap.Get("hash")
	log.Printf("try to login user: %v", name)
	var curator model.Curator
	err = connection.MySqlConnection.SelectOne(&curator, "select * from benutzer where Name=?", name)
	if err == nil && curator.ValidateHash(hash) {
		log.Printf("login successfull")
		object, _ := json.Marshal(curator)
		token := randToken()
		resp.Write([]byte(token))
		resp.Write([]byte("\n"))
		resp.Write(object)

		session := Session{}
		session.Id = curator.Id
		session.Name = curator.Name
		session.Time = time.Now()
		sessionMap[token] = session
		return
	}
	resp.WriteErrorString(http.StatusUnauthorized, err.Error())
}

func randToken() string {
	b := make([]byte, 40)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
