package statusresouce

import (
	"context"
)

func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {

	// TODO add version if not tracked
	// TODO when version changes set updating status
	// TODO when updating state is set and guest cluster is updated set updated status

	return nil
}
