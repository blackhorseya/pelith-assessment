package usecase

import (
	"context"
)

// Next is a function that represents the next step in the pipeline.
type Next func(c context.Context) error
