package models

import (
	"errors"

	"github.com/Izzat-Khudoyberganov/dictionary-app/db"
	"github.com/Izzat-Khudoyberganov/dictionary-app/utils"
)

type Admin struct {
	ID       int64
	Login    string `binding:"required"`
	Password string `binding:"required"`
}

func (a Admin) SaveAdmin() error {
	query := "INSERT INTO admin(login, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(a.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(a.Login, hashedPassword)

	if err != nil {
		return err
	}

	adminId, err := result.LastInsertId()

	a.ID = adminId

	return err
}

func (admin *Admin) ValidateAdmin() error {
	query := "SELECT id, password FROM admin WHERE login = ?"
	row := db.DB.QueryRow(query, admin.Login)

	var retrievedPassword string
	err := row.Scan(&admin.ID, &retrievedPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(admin.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
