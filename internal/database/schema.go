package database

import (
	"github.com/abigpotostew/go-api/internal/model"
	"gorm.io/gorm"
)

//type User struct {
//	ID           uint
//	Name         string
//	Email        *string
//	Age          uint8
//	Birthday     *time.Time
//	MemberNumber sql.NullString
//	ActivatedAt  sql.NullTime
//	CreatedAt    time.Time
//	UpdatedAt    time.Time
//}

// createSchema creates database schema for User and Story models.
func CreateSchema(db *gorm.DB) error {
	models := []interface{}{
		(*model.Recipe)(nil),
	}

	for _, model := range models {
		err := db.AutoMigrate(model)

		//err := db.Model(model).CreateTable(&orm.CreateTableOptions{
		//	Temp: true,
		//})
		if err != nil {
			return err
		}
	}
	return nil
}
