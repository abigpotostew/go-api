package db

import "github.com/uptrace/bun"

type ServiceDep struct {
	Database *bun.DB
}
