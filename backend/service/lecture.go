package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-app/dbutil"

	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
)

type Review struct {
	db *sqlx.DB
}

func NewReviewService(db *sqlx.DB) *Review {
	return &Review{db}
}

func (a *Review) Create(newReview *model.Review) (int64, error) {
	var createdID int64
	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateReview(tx, newReview)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		createdID = id
		return err
	}); err != nil {
		return 0, errors.Wrap(err, "failed comment insert transaction")
	}
	return createdID, nil
}
