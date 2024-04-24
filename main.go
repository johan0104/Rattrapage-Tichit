package main

import (
	"TICHIT/router"
	"TICHIT/template"
)

func main() {
	template.InitTemplate()
	router.Serveur()
}