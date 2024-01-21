package adapter

import (
	"context"
	"projects/LDmitryLD/repository/app/internal/infrastructure/db/tabler"
)

type SQLAdapterer interface {
	BuildSelect(tableName string, condition Condition, fields ...string) (string, []interface{}, error)
	Create(ctx context.Context, entity tabler.Tabler, ops ...interface{}) error
	Update(ctx context.Context, entity tabler.Tabler, condition Condition, ops ...interface{}) error
	Delete(ctx context.Context, tablename string, id int) error
	List(ctx context.Context, dest interface{}, table tabler.Tabler, condition Condition, opts ...interface{}) error
}
