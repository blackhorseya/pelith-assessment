package pg

import (
	"time"
)

const migrationFolder = "file://scripts/migrations"

const (
	defaultTimeout  = 5 * time.Second
	defaultLimit    = 10
	defaultMaxLimit = 100
)
