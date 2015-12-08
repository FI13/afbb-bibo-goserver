package model

type Curator struct {
	id           int    `db:"Id, primarykey, autoincrement" json:"id"`
	name         string `db:"Name,size:50" json:"name"`
	salt         string `db:"Salt,size:255" json:"salt"`
	passwordHash string `db:"Hash,size:255" json:"-"`
}
