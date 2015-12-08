package model

import (
	"database/sql"
)

type Borrower struct {
	id          int            `db:"Id" json:"id"`
	surname     string         `db:"Nachname" json:"surname"`
	forename    sql.NullString `db:"Vorname" json:"forename,omitempty"`
	info        sql.NullString `db:"Info" json:"info,omitempty"`
	email       sql.NullString `db:"EMail" json:"email,omitempty"`
	phoneNumber sql.NullString `db:"Telefon" json:"phoneNumber,omitempty"`
}
