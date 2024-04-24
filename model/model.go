package model

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "errors"
    "log"
)

type Article struct {
    ID          int    `json:"id"`
    Image       struct {
        URL string `json:"url"`
    } `json:"image"`
    Nom         string `json:"nom"`
    Suite       string `json:"suite"`
    Prix        string `json:"prix"`
    Reduc       string `json:"reduc"`
    Description string `json:"description"`
    Taille      string `json:"taille"`
}

type Articles struct {
    Articles []Article `json:"articles"`
}

var articlesData Articles

func LoadArticlesData(filePath string) (Articles, error) {
    var articles Articles
    jsonData, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Printf("Error reading JSON file: %s\n", err)
        return articles, err
    }
    err = json.Unmarshal(jsonData, &articles)
    if err != nil {
        fmt.Printf("Error decoding JSON data: %s\n", err)
        return articles, err
    }
    return articles, nil
}

func init() {
    var err error
    articlesData, err = LoadArticlesData("chemin/vers/le/fichier.json")
    if err != nil {
        log.Fatalf("Error loading articles data: %s", err)
    }
}

func GetProductByID(id int) (*Article, error) {
    for _, article := range articlesData.Articles {
        if article.ID == id {
            return &article, nil
        }
    }
    return nil, errors.New("product not found")
}

