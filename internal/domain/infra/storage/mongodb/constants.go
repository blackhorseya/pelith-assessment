package mongodb

import (
	"time"
)

const (
	dbName = "trading-ace"

	defaultTimeout = 5 * time.Second
	maxLimit       = int64(100)
	defaultLimit   = int64(10)
)
