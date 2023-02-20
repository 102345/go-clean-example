package userrepository

import (
	"context"
	"log"
)

func (repository repository) Delete(id uint64) error {
	ctx := context.Background()

	_, err := repository.db.Exec(
		ctx,
		"delete from user_api where id = $1",
		id,
	)

	if err != nil {
		log.Printf("Error delete in repository User: %s", err)
		return err
	}

	return nil
}
