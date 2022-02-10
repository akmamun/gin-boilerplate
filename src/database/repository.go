package database

import (
	"context"
)

func Save(ctx context.Context, model interface{}) error {
	err := _Save(ctx, model)
	return err
}
func Get(ctx context.Context, model interface{}, orderBy string, limit, offSet int) error {
	query := _Get(ctx, model, orderBy, limit, offSet)
	return query
}

func update() {

}
