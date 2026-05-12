package storage

import (
	"database/sql"
	"log"
	"my-api/models"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func GetAllMessages() ([]models.Request, error){
	rows, err := DB.Query("SELECT sender, content FROM messages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var reqs []models.Request

	for rows.Next() {
		var req models.Request
		if err := rows.Scan(&req.Sender, &req.Content); err != nil {
			return nil, err
		}
		reqs = append(reqs, req)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return reqs, nil
}

func InsertMessage(user, content string) {
	insertTableSQL := `INSERT INTO messages(sender, content) VALUES (?, ?)`
	_, err := DB.Exec(insertTableSQL, user, content)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("SUCCESS! Inserted message in DB")
}

func createTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS messages (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        sender TEXT,
        content TEXT
    );`
	if _, err := DB.Exec(createTableSQL); err != nil {
		log.Fatal("Failed to create table")
	}
}

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "database.db")
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	createTable()

	log.Println("SUCCESS! Connected to DB")

}