package productrepository

import (
	"context"
)

func (repository repository) Delete(id uint64) error {
	ctx := context.Background()

	_, err := repository.db.Exec(
		ctx,
		"delete from product where id = $1",
		id,
	)

	if err != nil {
		return err
	}

	return nil
}
