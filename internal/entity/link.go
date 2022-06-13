package entity

import "time"

type Link struct {
	ID               int64     `json:"id,omitempty"`
	Origin           string    `json:"origin" validation:"required"`
	RedirectLocation string    `json:"redirect_to,omitempty"`
	Hash             string    `json:"hash,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
}

//nolint
//lint:ignore U1000  tableName is being used by go-pg
type LinkDB struct {
	tableName struct{}  `pg:"links"`
	ID        int64     `pg:"id" db:"id"`
	Origin    string    `pg:"origin,notnull" db:"origin"`
	Hash      string    `pg:"hash,notnull" db:"link_hash"`
	UpdatedAt time.Time `pg:"updated_at,notnull" db:"updated_at"`
	CreatedAt time.Time `pg:"created_at,notnull" db:"created_at"`
}

func (LinkDB) TableName() string {
	return "urls"
}

func (u *Link) ToDB() *LinkDB {
	return &LinkDB{
		ID:        u.ID,
		Origin:    u.Origin,
		Hash:      u.Hash,
		UpdatedAt: u.CreatedAt,
		CreatedAt: u.UpdatedAt,
	}
}

func (db *LinkDB) ToWeb() *Link {
	return &Link{
		ID:        db.ID,
		Origin:    db.Origin,
		Hash:      db.Hash,
		UpdatedAt: db.CreatedAt,
		CreatedAt: db.UpdatedAt,
	}
}
