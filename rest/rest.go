package rest

import (
	"net/http"
	"strings"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

type GMError struct {
	Meta *Meta `json:"meta"`
    arg  int
    prob string
}

type Meta struct {
	Code   int64    `json:"code"`  
	Errors []string `json:"errors"`
}

func (e GMError) Error() string {
	if e.Meta != nil {
		return strings.Join(e.Meta.Errors, " + ")
	}
	return "unexpected response"
}