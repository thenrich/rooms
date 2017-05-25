package main

import (
	"google.golang.org/appengine"
	"golang.org/x/net/context"
	"net/http"
)

func AppEngineContext(r *http.Request) context.Context {
	return appengine.NewContext(r)
}

func init() {
	http.Handle("/", Router())

}
