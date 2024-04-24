package main

import (
	"TICHIT/router"
	initTemplate "TICHIT/template"
)

func main() {
	initTemplate.InitTemplate()
	router.Serveur()
}