package messages

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Message struct {
	ID      uint64 `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	DB      *sql.DB
}

func NewMessage() *Message {
	db, err := sql.Open("sqlite3", "./message.sql")
	if err != nil {
		log.Fatal(err)
	}

	return &Message{
		DB: db,
	}
}

func (m *Message) Add(message Message) {
	defer m.DB.Close()

	_, err := m.DB.Exec("INSERT INTO messages (title, content) VALUES (?,?)", message.Title, message.Content)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Message) All() ([]Message, error) {
	defer m.DB.Close()

	var messages []Message

	query, err := m.DB.Query("SELECT title, content FROM messages;")
	if err != nil {
		return nil, err
	}

	for query.Next() {
		var message Message
		err := query.Scan(&message.Title, &message.Content)
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}
	return messages, nil
}
