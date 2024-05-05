package db

import (
	"embed"
)

//go:embed migrations/*.sql
var Content embed.FS
