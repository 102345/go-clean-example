package productrepository

import (
	"context"
)

func (repository repository) Delete(id uint64) error {
	ctx := context.Background()

	err := repository.db.QueryRow(
		ctx,
		"delete from product where id = $1",
		id,
	)

	if err != nil {
		return ctx.Err()
	}

	return nil
}
