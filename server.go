package main

import (
	"github.com/TarsCloud/TarsGo/tars"
	"impl/web"
)

func main() {
	web.AddServant()
	tars.Run()
}
