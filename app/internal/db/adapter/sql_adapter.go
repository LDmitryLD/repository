package adapter

import (
	"context"
	"database/sql"
	"fmt"
	"projects/LDmitryLD/repository/app/internal/infrastructure/db/tabler"
	"reflect"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type SQLAdapter struct {
	db         *sqlx.DB
	sqlBuilder sq.StatementBuilderType
}

func NewSQLAdapter(db *sqlx.DB) *SQLAdapter {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return &SQLAdapter{
		db:         db,
		sqlBuilder: builder,
	}
}

func (s *SQLAdapter) BuildSelect(tableName string, condition Condition, fields ...string) (string, []interface{}, error) {
	if condition.ForUpdate {
		temp := []string{"FOR UPDATE"}
		temp = append(temp, fields...)
		fields = temp
	}

	queryRaw := s.sqlBuilder.Select(fields...).From(tableName)

	if condition.Equal != nil {
		for field, args := range condition.Equal {
			queryRaw = queryRaw.Where(sq.Eq{field: args})
		}
	}

	queryRaw = queryRaw.Where(sq.Eq{"deleted_at": sql.NullString{}})

	if condition.NotEqual != nil {
		for field, args := range condition.NotEqual {
			queryRaw = queryRaw.Where(sq.NotEq{field: args})
		}
	}

	if condition.Order != nil {
		for _, order := range condition.Order {
			direction := "DESC"
			if order.Asc {
				direction = "ASC"
			}
			queryRaw = queryRaw.OrderBy(fmt.Sprintf("%s %s", order.Field, direction))
		}
	}

	if condition.LimitOffset != nil {
		if condition.LimitOffset.Limit > 0 {
			queryRaw.Limit(uint64(condition.LimitOffset.Limit))
		}
		if condition.LimitOffset.Offset > 0 {
			queryRaw.Offset(uint64(condition.LimitOffset.Offset))
		}
	}

	return queryRaw.ToSql()
}

func (s *SQLAdapter) Create(ctx context.Context, entity tabler.Tabler, ops ...interface{}) error {
	info := tabler.GetStructInfo(entity, filterByTag("db_ops", "create"))
	m := prepareMap(info)

	queryInsert := s.sqlBuilder.Insert(entity.TableName()).SetMap(m)

	query, args, err := queryInsert.ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, query, args...)

	return err
}

func (s *SQLAdapter) Update(ctx context.Context, entity tabler.Tabler, condition Condition, ops ...interface{}) error {
	info := tabler.GetStructInfo(entity, filterByTag("db_ops", "update"))

	m := prepareMap(info)

	queryUpdate := s.sqlBuilder.Update(entity.TableName())

	if condition.Equal != nil {
		for field, args := range condition.Equal {
			queryUpdate = queryUpdate.Where(sq.Eq{field: args})
		}
	}

	if condition.NotEqual != nil {
		for field, args := range condition.NotEqual {
			queryUpdate = queryUpdate.Where(sq.NotEq{field: args})
		}
	}

	queryUpdate = queryUpdate.SetMap(m)

	query, args, err := queryUpdate.ToSql()
	if err != nil {
		return err
	}

	res, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()

	return err
}

func (s *SQLAdapter) Delete(ctx context.Context, tablename string, id int) error {

	queryDelete := s.sqlBuilder.Update(tablename).Set("deleted_at", time.Now().String()).Where(sq.Eq{"id": id})

	query, args, err := queryDelete.ToSql()
	if err != nil {
		return err
	}

	res, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()

	return err
}

func (s *SQLAdapter) List(ctx context.Context, dest interface{}, table tabler.Tabler, condition Condition, opts ...interface{}) error {
	var ops []func(*[]reflect.StructField)
	for _, opt := range opts {
		if option, ok := opt.(func(*[]reflect.StructField)); ok {
			ops = append(ops, option)
		}
	}

	info := tabler.GetStructInfo(table, ops...)

	query, args, err := s.BuildSelect(table.TableName(), condition, info.Fields...)
	if err != nil {
		return err
	}

	err = s.db.SelectContext(ctx, dest, query, args...)

	return err
}
