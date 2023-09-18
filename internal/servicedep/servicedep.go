package servicedep

import (
	"github.com/abigpotostew/go-api/internal/database"
	"github.com/abigpotostew/go-api/internal/servicedep/db"
	"os"
)

type ServiceDep struct {
	DB *db.ServiceDep
}

func NewServiceDep() *ServiceDep {
	var dsn = os.Getenv("DATABASE_URL")
	return &ServiceDep{
		DB: &db.ServiceDep{
			Database: database.ConnectDatabase(dsn),
		},
	}
}
