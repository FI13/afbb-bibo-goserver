package model

import (
	"database/sql"
	"encoding/json"
)

type Borrower struct {
	Id          int            `db:"Id"`
	Surname     string         `db:"Nachname"`
	Forename    sql.NullString `db:"Vorname"`
	Info        sql.NullString `db:"Info"`
	Email       sql.NullString `db:"EMail"`
	PhoneNumber sql.NullString `db:"Telefon"`
}

func (b *Borrower) MarshalJSON() ([]byte, error) {
	str := struct {
		Id          int    `json:"id"`
		Surname     string `json:"surname"`
		Forename    string `json:"forename,omitempty"`
		Info        string `json:"info,omitempty"`
		Email       string `json:"email,omitempty"`
		PhoneNumber string `json:"phoneNumber,omitempty"`
	}{b.Id, b.Surname, b.Forename.String, b.Info.String, b.Email.String, b.PhoneNumber.String}
	return json.Marshal(str)
}
