package models

type Article struct {
ID int `json:"id`
Title string `json:"title"`
Content string `json:"content"`
Tags []string `json:"tags"`
CreatedAt string `json:"created_at"`
UpdatedAt string `json:"updated_at"`
}
