package repository

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"rest-api/database"
	"time"
)

type Token struct {
	Id        int
	UserId    int
	Token     string
	ExpiredAt string
	CreatedAt string
}

func GenerateToken(length int) string {
	src := make([]byte, length)
	if _, err := rand.Read(src); err != nil {
		return ""
	}
	return hex.EncodeToString(src)
}

func (t Token) Create() Token {
	now := time.Now()
	t.Token = GenerateToken(20)
	t.CreatedAt = now.Format("2006-01-02 15:04:05")
	if t.ExpiredAt == "" {
		t.ExpiredAt = now.AddDate(0, 0, 1).Format("2006-01-02 15:04:05")
	}
	r, err := database.Instance().Exec(`INSERT INTO tokens (user_id, token, expired_at, created_at) VALUES (?,?,?,?)`, t.UserId, t.Token, t.ExpiredAt, t.CreatedAt)
	if err != nil {
		log.Println(err.Error())
	}
	id, _ := r.LastInsertId()
	t.Id = int(id)
	return t
}

func (t Token) Get() Token {
	rows, err := database.Instance().Query(`SELECT * FROM tokens WHERE token = ?`, t.Token)
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&t.Id, &t.UserId, &t.Token, &t.ExpiredAt, &t.CreatedAt)
		if err != nil {
			panic(err)
		}
	}
	return t
}
