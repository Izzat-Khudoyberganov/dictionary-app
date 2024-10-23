package models

import (
	"fmt"

	"github.com/Izzat-Khudoyberganov/dictionary-app/db"
)

type Phrasa struct {
	ID        int64
	Word      string `binding:"required"`
	Translate string `binding:"required"`
}

func (p Phrasa) SavePhrasa() error {
	query := "INSERT INTO phrasa(word, translate) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(p.Word, p.Translate)

	if err != nil {
		return err
	}

	phrasaId, err := result.LastInsertId()
	p.ID = phrasaId

	return err
}

func GetAllPhrasa() ([]Phrasa, error) {
	query := "SELECT id, word, translate FROM phrasa"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var phrasa []Phrasa

	for rows.Next() {
		var ph Phrasa

		err := rows.Scan(&ph.ID, &ph.Word, &ph.Translate)
		if err != nil {
			return nil, err
		}

		phrasa = append(phrasa, ph)
	}

	return phrasa, nil
}

func (p Phrasa) UpdatePhrasa() error {
	query := "UPDATE phrasa SET word = ?, translate = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(p.Word, p.Translate)

	if err != nil {
		return fmt.Errorf("failed to execute update: %w", err)
	}
	return nil
}

func GetPhrasaById(id int64) (*Phrasa, error) {
	query := "SELECT id, word, translate FROM phrasa WHERE id = ?"

	row := db.DB.QueryRow(query, id)
	var phrasa Phrasa
	err := row.Scan(&phrasa.ID, &phrasa.Word, &phrasa.Translate)
	if err != nil {
		return nil, err
	}

	return &phrasa, nil
}

func (p Phrasa) DeletePhrasa() error {
	query := "DELETE FROM phrasa WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(p.ID)

	return err
}
