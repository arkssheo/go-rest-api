package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password123"
	dbname   = "articlesdb"
)

var psqlInfo string
var dbConn *sql.DB

func initDb() {
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	dbConn = db

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	newArticle := Article{
		Title:   "Test Article Insert",
		Desc:    "A new description",
		Content: "My brand new and longer content column, yay!",
	}

	addedID, err := insertArticle(&newArticle)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully inserted! Id:", addedID)
}

func insertArticle(article *Article) (int, error) {
	if dbConn == nil {
		return -1, errors.New("DB is nil")
	}

	if article == nil {
		return -1, errors.New("Article is nil")
	}
	sqlStatement := `
		INSERT INTO article (title, description, content)
		VALUES ($1, $2, $3)
		RETURNING id`
	id := 0
	err := dbConn.QueryRow(sqlStatement, article.Title, article.Desc, article.Content).Scan(&id)
	if err != nil {
		return -1, err
	}

	fmt.Println("New record ID is:", id)
	return id, nil
}
