package models

import (
	"example.com/BookEvent/utils"

	"example.com/BookEvent/db"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.Id = userId
	return nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.Id, &retrievedPassword)

	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPassword(u.Password, retrievedPassword)

	if !passwordIsValid {
		return err
	}

	return nil
}

func GetAllUsers() ([]*User, error) {
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func GetUserById(id int64) (*User, error) {
	query := "SELECT * FROM users WHERE id =?"
	row := db.DB.QueryRow(query, id)

	var user User
	err := row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) Deletes() error {
	query := "DELETE FROM users WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Id)
	return err
}
