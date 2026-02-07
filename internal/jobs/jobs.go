package jobs

import "context"

// Job defines a unit of work for background processing.
type Job interface {
	Name() string
	Run(ctx context.Context) error
}
