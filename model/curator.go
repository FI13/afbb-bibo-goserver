package model

type Curator struct {
	Id           int    `db:"Id, primarykey, autoincrement" json:"id"`
	Name         string `db:"Name,size:50" json:"name"`
	Salt         string `db:"Salt,size:255" json:"salt"`
	PasswordHash string `db:"Hash,size:255" json:"-"`
}

func (c Curator) ValidateHash(hash string) bool {
	return hash == c.PasswordHash
}
