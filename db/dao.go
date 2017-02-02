package db

import (
	"database/sql"
	"errors"
	"github.com/adorsys/golang-chi-rest-db-oauth-sample/model"
	_ "github.com/lib/pq"
	"log"
)

var connection *sql.DB

func init() {
	con, err := sql.Open("postgres", "user=go dbname=go sslmode=disable")
	if err != nil {
		log.Fatal("Could not open DB: ", err)
	}

	err = con.Ping()
	if err != nil {
		log.Fatal("Could not open DB: ", err)
	}

	log.Println("DB initialized")
	connection = con
}

func GetArticle(id int) (*model.Article, error) {
	result := model.Article{}
	err := connection.QueryRow("SELECT id, title FROM article WHERE id = $1", id).Scan(&result.ID, &result.Title)

	if err != nil {
		return nil, errors.New("Article not found")
	}

	return &result, nil
}

func ListArticles() ([]*model.Article, error) {
	rows, err := connection.Query("SELECT id, title FROM article")
	if err != nil {
		return nil, errors.New("No articles found")
	}

	result := []*model.Article{}

	for rows.Next() {
		cur := model.Article{}
		err := rows.Scan(&cur.ID, &cur.Title)
		if err != nil {
			return nil, errors.New("Could not map articles")
		}
		result = append(result, &cur)
	}

	return result, nil
}

func CreateArticle(title string) (*model.Article, error) {
	result := model.Article{}

	if err := connection.QueryRow("INSERT INTO article(title) VALUES ($1) RETURNING id, title", title).Scan(
		&result.ID, &result.Title); err != nil {
		log.Print("Could not insert article: ", err)
		return nil, errors.New("Could not insert article")
	} else {
		return &result, nil

	}
}

func DeleteArticle(id int64) error {
	_, err := connection.Exec("DELETE FROM article WHERE article.id = $1", id)
	return err
}
