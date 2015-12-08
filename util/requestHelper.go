package util

import (
	"github.com/emicklei/go-restful"
	"net/url"
)

func ExtractParameters(req *restful.Request) (map[string][]string, error) {
	m, err := url.ParseQuery(req.Request.URL.RawQuery)
	return m, err
}
