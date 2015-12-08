package util

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
)

func ServiceResponse(response *restful.Response, object interface{}, err error) {
	log.Printf("response %s", err)
	if err == nil {
		response.WriteEntity(object)
		return
	}
	response.WriteErrorString(http.StatusNotFound, err.Error())
}
