package config

import (
	"net/http"
)

type Ambassador struct {
	Ambassador string
	Path       string
	URL        string
	HTTPVerb   string
}

var HTTPVerbFunctionMap = map[string]interface{}{
	"GET":  http.Get,
	"POST": http.Post,
}
