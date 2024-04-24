package controller

import (
	"TICHIT/model"
	initTemplate "TICHIT/template"
	"net/http"
    "path/filepath"
    "os"
    "fmt"
    "strconv"
    "strings"
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

func DetailHandler(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Path[len("/detail/"):] // Assurez-vous que votre route correspond à ce format
    fmt.Println("Received ID:", idStr) // Debug: Afficher l'ID reçu pour vérifier
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

    product, err := model.GetProductByID(id) 
    if err != nil {
        http.Error(w, "Product not found", http.StatusNotFound)
        return
    }

    initTemplate.Temp.ExecuteTemplate(w, "detail", product)
}

func AjoutHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        err := r.ParseForm()
        if err != nil {
            http.Error(w, "Error parsing form", http.StatusBadRequest)
            return
        }
        prix := r.FormValue("prix")
        prix = strings.TrimSpace(prix)
        if _, err := strconv.ParseFloat(prix, 64); err != nil {
            http.Error(w, "Invalid price format", http.StatusBadRequest)
            return
        }

        article := model.Article{
            Nom:         r.FormValue("nom"),
            Suite:       r.FormValue("suite"),
            Prix:        prix,
            Reduc:       r.FormValue("reduc"),
            Description: r.FormValue("description"),
            Taille:      r.FormValue("taille"),
        }

        if err := model.AddArticle(&article); err != nil {
            http.Error(w, "Failed to add article", http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/accueil", http.StatusSeeOther)
    } else {
        // Afficher le formulaire d'ajout
        initTemplate.Temp.ExecuteTemplate(w, "ajout", nil)
    }
}

