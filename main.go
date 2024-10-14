package main

import(
	"log"
	"net/http"
	"github.com/JAANVI999/blog-api/internal/handlers"
	"github.com/JAANVI999/blog-api/internal/repository"
)
func main(){

	repository.InitDB()
	//define the routes
	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request){
		switch r.Method{
			case http.MethodPost:
				handlers.CreateArticle(w,r)
			case http.MethodPut:
                handlers.UpdateArticle(w,r)
			case http.MethodGet:
				handlers.GetArticle(w,r)
			case http.MethodDelete:
				handlers.DeleteArticle(w,r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	log.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))

}