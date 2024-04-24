package controller

import (
	"TICHIT/model"
	initTemplate "TICHIT/template"
	"net/http"
)

func AccueilHandler(w http.ResponseWriter, r *http.Request) {
    articles, err := model.LoadArticlesData("../data/articles.json")
    if err != nil {
        http.Error(w, "Unable to load articles data", http.StatusInternalServerError)
        return
    }
    initTemplate.Temp.ExecuteTemplate(w, "accueil", articles)
}

func AjoutHandler(w http.ResponseWriter, r *http.Request) {

	initTemplate.Temp.ExecuteTemplate(w, "ajout", nil)
}	

func DetailHandler(w http.ResponseWriter, r *http.Request) {

	initTemplate.Temp.ExecuteTemplate(w, "detail", nil)
}	
