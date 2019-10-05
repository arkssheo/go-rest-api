package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// fmt.Fprintf(w, "Key: "+key)
	for _, article := range articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var deleted Article
	for index, article := range articles {
		if article.Id == id {
			deleted = articles[index]
			articles = append(articles[:index], articles[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(deleted)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	// grab the id from the route param, like in delete
	vars := mux.Vars(r)
	id := vars["id"]

	// read the body for the data to update, create the new article to replace
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newArticle Article

	// create the article by unmarshall and passing the obj reference
	json.Unmarshal(reqBody, &newArticle)

	// Same as delete, run through articles and find the one to update
	for index, article := range articles {
		if article.Id == id {
			// once found, replace the article with the newly created
			articles[index] = newArticle
		}
	}

	// end with returning all the articles
	json.NewEncoder(w).Encode(articles)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewArticle")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	articles = append(articles, article)
	json.NewEncoder(w).Encode(article)
	//fmt.Fprintf(w, "%+v", string(reqBody))
}
