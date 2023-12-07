package repository

import (
	"fmt"

	ent2 "github.com/42milez/go-oidc-server/app/pkg/ent/ent"
)

var errEntNotFoundError = &ent2.NotFoundError{}

func rollback(tx *ent2.Tx, err error) error {
	if retErr := tx.Rollback(); retErr != nil {
		return fmt.Errorf("%w: %v", err, retErr)
	}
	return err
}
