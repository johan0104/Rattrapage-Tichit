package model

import (
    "encoding/json"
    "io/ioutil"
)

type Article struct {
    ID          int    `json:"id"`
    Image       struct {
        URL string `json:"url"`
    } `json:"image"`
    Nom         string `json:"nom"`
    Prix        string `json:"prix"`
    Reduc       string `json:"reduc"`
    Description string `json:"description"`
    Taille      string `json:"taille"`
}

type Articles struct {
    Articles []Article `json:"articles"`
}

func LoadArticlesData(filePath string) (Articles, error) {
    var articles Articles
    jsonData, err := ioutil.ReadFile(filePath)
    if err != nil {
        return articles, err
    }
    err = json.Unmarshal(jsonData, &articles)
    return articles, err
}
