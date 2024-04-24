package template

import (
	"fmt"
	"html/template"
)

var Temp *template.Template

func InitTemplate() {
	temp, errTemp := template.ParseGlob("./template/*.html")
	if errTemp != nil {
		// Utilisation de fmt.Printf pour le formatage correct
		fmt.Printf("Erreur => %s\n", errTemp.Error())
	}
	Temp = temp
}