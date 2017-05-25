package main

import (
	"golang.org/x/net/context"
	"github.com/thenrich/rooms/conf"

)

type AppContext struct {
	Ctx          context.Context
	TwilioConfig *conf.TwilioConfig
}
