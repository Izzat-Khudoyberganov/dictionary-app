package models

import (
	"fmt"

	"github.com/Izzat-Khudoyberganov/dictionary-app/db"
)

type Dictionary struct {
	ID          int64
	Word        string `binding:"required"`
	Translate   string `binding:"required"`
	Description string `binding:"required"`
	Example     string `binding:"required"`
}

func (d Dictionary) SaveDictionary() error {
	query := "INSERT INTO dictionary(word, translate, description, example) VALUES (?, ?, ?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(d.Word, d.Translate, d.Description, d.Example)

	if err != nil {
		return err
	}

	dictionaryId, err := result.LastInsertId()
	d.ID = dictionaryId

	return err
}

func GetAllDictionary() ([]Dictionary, error) {
	query := "SELECT id, word, translate, description, example FROM dictionary"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var dictionary []Dictionary

	for rows.Next() {
		var dic Dictionary

		err := rows.Scan(&dic.ID, &dic.Word, &dic.Translate, &dic.Description, &dic.Example)
		if err != nil {
			return nil, err
		}

		dictionary = append(dictionary, dic)
	}

	return dictionary, nil
}

func (d Dictionary) UpdateDictionary() error {
	query := "UPDATE dictionary SET word = ?, translate = ?, description = ?, example = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(d.Word, d.Translate, d.Description, d.Example)

	if err != nil {
		return fmt.Errorf("failed to execute update: %w", err)
	}
	return nil
}

func (d Dictionary) DeleteDictionary() error {
	query := "DELETE FROM dictionary WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(d.ID)

	return err
}

func GetAllDictionaryById(id int64) (*Dictionary, error) {
	query := "SELECT id, word, translate, description, example FROM dictionary WHERE id = ?"

	row := db.DB.QueryRow(query, id)
	var dictionary Dictionary
	err := row.Scan(&dictionary.ID, &dictionary.Word, &dictionary.Translate, &dictionary.Description, &dictionary.Example)
	if err != nil {
		return nil, err
	}

	return &dictionary, nil
}
