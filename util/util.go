package util

import (
	"encoding/json"
	"log"
)

func LogAsJson(object interface{}) {
	test, _ := json.Marshal(object)
	log.Printf(string(test))
}
