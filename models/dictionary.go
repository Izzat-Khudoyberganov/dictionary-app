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

// Save function
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

// Get all dictionary with pagination
func GetDictionary(page, limit int) ([]Dictionary, int, error) {
	offset := (page - 1) * limit

	query := "SELECT id, word, translate, description, example FROM dictionary LIMIT ? OFFSET ?"
	totalCountQuery := "SELECT COUNT(*) FROM dictionary"

	var totalDictionary int
	err := db.DB.QueryRow(totalCountQuery).Scan(&totalDictionary)
	if err != nil {
		return nil, 0, err
	}

	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var dictionary []Dictionary

	for rows.Next() {
		var dic Dictionary
		rows.Scan(&dic.ID, &dic.Word, &dic.Translate, &dic.Description, &dic.Example)
		if err != nil {
			return nil, 0, err
		}
		dictionary = append(dictionary, dic)
	}

	return dictionary, totalDictionary, nil
}

// Update dictionary
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

// Delete dictionary
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

func GetDictionaryById(id int64) (*Dictionary, error) {
	query := "SELECT id, word, translate, description, example FROM dictionary WHERE id = ?"

	row := db.DB.QueryRow(query, id)
	var dictionary Dictionary
	err := row.Scan(&dictionary.ID, &dictionary.Word, &dictionary.Translate, &dictionary.Description, &dictionary.Example)
	if err != nil {
		return nil, err
	}

	return &dictionary, nil
}
