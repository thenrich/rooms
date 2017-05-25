package main

import (
	_ "github.com/thenrich/rooms/conf"
	_ "github.com/thenrich/rooms/logger"
	_ "github.com/thenrich/rooms/util"
	"google.golang.org/appengine"
)

func main() {
	appengine.Main()
}
