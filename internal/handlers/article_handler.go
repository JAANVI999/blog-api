package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JAANVI999/blog-api/internal/models"
	"github.com/JAANVI999/blog-api/internal/repository"
)
func CreateArticle(w http.ResponseWriter, r *http.Request){
	var article models.Article
	err:= json.NewDecoder(r.Body).Decode(&article)
	if err !=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = repository.CreateArticle(article)
	if err!= nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article)
}

func GetArticle(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	if id == ""{
		http.Error(w, "Missing article ID", http.StatusBadRequest)
		return
	}
	articleID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}
	article, err := repository.GetArticle(articleID)
	if err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(article)
}
func UpdateArticle(w http.ResponseWriter, r *http.Request){
	var article models.Article
	err:= json.NewDecoder(r.Body).Decode(&article)
	if err !=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}
	err = repository.UpdateArticle(article)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
func DeleteArticle(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	if id == ""{
		http.Error(w, "Missing article ID", http.StatusBadRequest)
		return
	}
	articleID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}
	err = repository.DeleteArticle(articleID)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}