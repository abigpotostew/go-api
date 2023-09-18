package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Recipe struct {
	ID   int64  `bun:"id,pk,autoincrement"`
	Name string `bun:"name,notnull"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt time.Time `bun:",soft_delete,nullzero"` //Enable soft deletes on the model.
}
type User struct {
	bun.BaseModel `bun:"table:recipes"`

	Id   int64 `bun:",pk,autoincrement"`
	Name string
}
