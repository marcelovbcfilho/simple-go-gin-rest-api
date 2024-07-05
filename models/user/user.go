package user

import (
	"database/sql"
	"simple-go-gin-rest-api/infrastructure"
	"simple-go-gin-rest-api/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

var insertQuery = `INSERT INTO users (email, password) VALUES (?, ?)`

func (u *User) Save() error {
	stmt, err := infrastructure.DB.Prepare(insertQuery)
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err)
		}
	}(stmt)

	password, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.Id = id
	return nil
}
