package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/JAANVI999/blog-api/internal/models"
)

func CreateArticle(article models.Article) error{
	tags := strings.Join(article.Tags, ",") 
	query := `INSERT INTO articles (title, content, tags, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id`
	var id int
	err:= DB.QueryRow(query, article.Title, article.Content, tags).Scan(&id)
	if err != nil{
		return err
	}
    article.ID = id
	fmt.Println("Newly created article ID: ", id)
	return nil
}

func GetArticle(id int) (models.Article, error){
	var article models.Article
	var tags string
	query := `SELECT * FROM articles WHERE id = $1`
    err:= DB.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Content, &tags, &article.CreatedAt, &article.UpdatedAt)
	if err != nil{
		if err == sql.ErrNoRows{
			return article, fmt.Errorf("Article not found")
		}
		return article, err
	}
	article.Tags = strings.Split(tags, ",")
	return article, nil
}
func UpdateArticle(article models.Article) error{
	tags := strings.Join(article.Tags, ",") 
	query:= `UPDATE articles SET title =$1, content=$2, tags=$3, updated_at=NOW() WHERE id=$4`
	_, err:= DB.Exec(query, article.Title, article.Content, tags, article.ID)
	if err != nil{
		return err
	}
	fmt.Println("Article updated successfully")
	return nil
}
func DeleteArticle(id int) error{
	query:= `DELETE FROM articles WHERE id=$1`
	_, err:= DB.Exec(query, id)
	if err!= nil{
		return err
	}
	fmt.Println("Article deleted successfully")
	return nil
}
