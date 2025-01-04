package controllers

import (
	"database/sql"

	"github.com/Alfazal007/ctr_solana/internal/database"
	"github.com/cloudinary/cloudinary-go/v2"
)

type ApiConf struct {
	DB         *database.Queries
	Cloudinary *cloudinary.Cloudinary
	SQLDB      *sql.DB
}
