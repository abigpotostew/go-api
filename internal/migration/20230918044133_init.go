package migration

import (
	"context"
	"fmt"
	"github.com/abigpotostew/go-api/internal/model"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Println(" [up migration] for recipe model")
		res, err := db.NewCreateTable().
			Model((*model.Recipe)(nil)).
			Exec(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Printf(" [up migration] for recipe model: %v\n", res)
		return err

	}, func(ctx context.Context, db *bun.DB) error {
		res, err := db.NewDropTable().Model((*model.Recipe)(nil)).Exec(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Printf(" [down migration] for recipe model: %v\n", res)
		return err
	})
}
