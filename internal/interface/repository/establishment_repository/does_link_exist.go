package establishment_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
)

func (e *establishmentRepository) DoesLinkExist(ctx context.Context, link string) (bool, error) {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return false, errors.New("db not found in context")
	}

	var cnt int64
	err := db.QueryRow(ctx, "select count(id) from establishments where link = $1", link).Scan(&cnt)
	if err != nil {
		return false, err
	}

	return cnt > 0, nil
}
