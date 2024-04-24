package controller

import (
	"TICHIT/model"
	initTemplate "TICHIT/template"
	"net/http"
    "path/filepath"
    "os"
    "fmt"
    "strconv"
)

func AccueilHandler(w http.ResponseWriter, r *http.Request) {
    cwd, _ := os.Getwd()
    filePath := filepath.Join(cwd, "/data/articles.json")
    fmt.Println("Attempting to load JSON from:", filePath)  // Afficher le chemin complet

    articles, err := model.LoadArticlesData(filePath)
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
    idStr := r.URL.Path[len("/detail/"):] // Assurez-vous que votre route correspond à ce format
    fmt.Println("Received ID:", idStr) // Debug: Afficher l'ID reçu pour vérifier
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

    product, err := GetProductByID(id) 
    if err != nil {
        http.Error(w, "Product not found", http.StatusNotFound)
        return
    }

    initTemplate.Temp.ExecuteTemplate(w, "detail", product)
}

