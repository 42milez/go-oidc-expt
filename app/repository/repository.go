package repository

import (
	"fmt"

	"github.com/42milez/go-oidc-server/app/ent/ent"
)

func rollback(tx *ent.Tx, err error) error {
	if retErr := tx.Rollback(); retErr != nil {
		return fmt.Errorf("%w: %v", err, retErr)
	}
	return err
}
