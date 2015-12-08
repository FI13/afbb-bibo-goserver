package util

import (
	"net/http"
	"net/url"

	"github.com/emicklei/go-restful"
)

type ParameterMap map[string]string

func ExtractParameters(req *restful.Request) (http.Header, error) {
	m, err := url.ParseQuery(req.Request.URL.RawQuery)
	for key, value := range m {
		for i := range value {
			req.Request.Header.Add(key, value[i])
		}
	}
	return req.Request.Header, err
}
