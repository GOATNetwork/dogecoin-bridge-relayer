package orm

import (
	"context"
	"fmt"
	"reflect"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// Repository is a standard struct for all repository structs.
type Repository[T any] struct {
	db    *gorm.DB
	model *T
}

func New[T any](db *gorm.DB, model T) *Repository[T] {
	return &Repository[T]{
		db:    db,
		model: &model,
	}
}

func (r *Repository[T]) Use(plugin gorm.Plugin) error {
	err := r.db.Use(plugin)
	if err != nil {
		return fmt.Errorf("orm.Repository.Use: %w", err)
	}

	return nil
}

// Session create new db session.
func (r *Repository[T]) Session(sess *gorm.Session) *Repository[T] {
	return &Repository[T]{db: r.db.Session(sess), model: r.model}
}

func (r *Repository[T]) WithContext(ctx context.Context) *Repository[T] {
	return &Repository[T]{db: r.db.WithContext(ctx), model: r.model}
}

// Debug start debug mode.
func (r *Repository[T]) Debug() *Repository[T] {
	return &Repository[T]{db: r.db.Debug(), model: r.model}
}

// Begin begins a transaction with any transaction options opts.
func (r *Repository[T]) Begin() *Repository[T] {
	return &Repository[T]{db: r.db.Begin(), model: r.model}
}

// Rollback rollbacks the changes in a transaction.
func (r *Repository[T]) Rollback() *gorm.DB {
	return r.db.Rollback()
}

// Commit commits the changes in a transaction.
func (r *Repository[T]) Commit() *gorm.DB {
	return r.db.Commit()
}

// Transaction start a transaction as a block, return error will roll back, otherwise to commit.
func (r *Repository[T]) Transaction(fn func(repo *Repository[T]) error) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		return fn(&Repository[T]{db: tx, model: r.model})
	})
	if err != nil {
		return fmt.Errorf("orm.Repository.Transaction: %w", err)
	}

	return nil
}

// Set store value with key into current db instance's context.
func (r *Repository[T]) Set(k string, v any) *Repository[T] {
	return &Repository[T]{db: r.db.Set(k, v), model: r.model}
}

// Get value with key from current db instance's context.
func (r *Repository[T]) Get(k string) (any, bool) {
	return r.db.Get(k)
}

// Insert inserts value, returning the inserted data's primary key in value's id.
func (r *Repository[T]) Insert(entity *T, opts ...OptionFunc) *gorm.DB {
	db := r.GetDB(opts...)

	return db.Model(entity).Create(entity)
}

func (r *Repository[T]) InsertMany(entities []*T, opts ...OptionFunc) *gorm.DB {
	db := r.GetDB(opts...)

	if len(entities) == 0 {
		return db
	}

	return db.Create(entities)
}

// InsertInBatches inserts value in batches of size.
func (r *Repository[T]) InsertInBatches(entities []*T, size int, opts ...OptionFunc) *gorm.DB {
	db := r.GetDB(opts...)

	if len(entities) == 0 {
		return db
	}

	return db.CreateInBatches(entities, size)
}

// Save updates value in database. If value doesn't contain a matching primary key, value is inserted.
func (r *Repository[T]) Save(entity *T, opts ...OptionFunc) *gorm.DB {
	db := r.GetDB(opts...)

	return db.Save(entity)
}

// UpdateOne updates attributes using callbacks. values must be a struct.
func (r *Repository[T]) UpdateOne(entity *T, opts ...OptionFunc) *gorm.DB {
	db := r.GetDB(opts...)

	return db.Model(entity).Updates(entity)
}

// UpdateOneWithMap updates attributes using callbacks. values must be a map.
func (r *Repository[T]) UpdateOneWithMap(entity *T, m map[string]any, opts ...OptionFunc) *gorm.DB {
	db := r.GetDB(opts...)

	return db.Model(entity).Updates(m)
}

func (r *Repository[T]) UpdatesByCondition(
	cond func(*gorm.DB) *gorm.DB, model *T, opts ...OptionFunc,
) *gorm.DB {
	db := r.GetDB(opts...)

	return db.Model(r.model).Scopes(cond).Updates(model)
}

func (r *Repository[T]) UpdatesByConditionWithMap(
	cond func(*gorm.DB) *gorm.DB, m map[string]any, opts ...OptionFunc,
) *gorm.DB {
	db := r.GetDB(opts...)

	return db.Model(r.model).Scopes(cond).Updates(m)
}

// Delete deletes value matching given conditions.
func (r *Repository[T]) Delete(entity *T, opts ...OptionFunc) *gorm.DB {
	db := r.GetDB(opts...)

	return db.Delete(entity)
}

func (r *Repository[T]) DeleteByID(id int64, opts ...OptionFunc) *gorm.DB {
	db := r.GetDB(opts...)

	return db.Delete(r.model, id)
}

func (r *Repository[T]) DeleteByCondition(cond func(*gorm.DB) *gorm.DB, opts ...OptionFunc) *gorm.DB {
	db := r.GetDB(opts...)

	var entity T

	return db.Scopes(cond).Delete(&entity)
}

// First finds the first record ordered by primary key, matching given conditions.
func (r *Repository[T]) First(opts ...OptionFunc) (*T, *gorm.DB) {
	db := r.GetDB(opts...)

	var model T

	db = db.First(&model)

	return &model, db
}

func (r *Repository[T]) Take(opts ...OptionFunc) (*T, *gorm.DB) {
	db := r.GetDB(opts...)

	var model T

	db = db.Take(&model)

	return &model, db
}

func (r *Repository[T]) Last(opts ...OptionFunc) (*T, *gorm.DB) {
	db := r.GetDB(opts...)

	var model T

	db = db.Last(&model)

	return &model, db
}

func (r *Repository[T]) FirstByID(id int64, opts ...OptionFunc) (*T, *gorm.DB) {
	db := r.GetDB(opts...)

	var entity T

	db = db.First(&entity, id)

	return &entity, db
}

func (r *Repository[T]) FirstByCondition(cond func(*gorm.DB) *gorm.DB, opts ...OptionFunc) (*T, *gorm.DB) {
	db := r.GetDB(opts...)

	var entity T

	db = db.Scopes(cond).First(&entity)

	return &entity, db
}

func (r *Repository[T]) FirstByConditionWithLock(
	cond func(*gorm.DB) *gorm.DB, opts ...OptionFunc,
) (*T, *gorm.DB) {
	db := r.GetDB(opts...)

	var entity T

	db = db.Clauses(clause.Locking{Strength: "UPDATE"}).Scopes(cond).First(&entity)

	return &entity, db
}

// LastByCondition finds the last record ordered by primary key, matching given conditions.
func (r *Repository[T]) LastByCondition(cond func(*gorm.DB) *gorm.DB, opts ...OptionFunc) (*T, *gorm.DB) {
	db := r.GetDB(opts...)

	var entity T

	db = db.Scopes(cond).Last(&entity)

	return &entity, db
}

// TakeByCondition finds the first record returned by the database in no specified order, matching given conditions.
func (r *Repository[T]) TakeByCondition(cond func(*gorm.DB) *gorm.DB, opts ...OptionFunc) (*T, *gorm.DB) {
	db := r.GetDB(opts...)

	var entity T

	db = db.Scopes(cond).Take(&entity)

	return &entity, db
}

func (r *Repository[T]) List(opts ...OptionFunc) (*Page[T], *gorm.DB) {
	return r.Paginate(nil, opts...)
}

func (r *Repository[T]) ListByIDs(ids []int64, opts ...OptionFunc) (*Page[T], *gorm.DB) {
	return r.PageByCondition(func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id IN (?)", ids)
	}, nil, opts...)
}

func (r *Repository[T]) ListByCondition(
	cond func(*gorm.DB) *gorm.DB, opts ...OptionFunc,
) (*Page[T], *gorm.DB) {
	return r.PageByCondition(cond, nil, opts...)
}

func (r *Repository[T]) ListByClause(
	where []clause.Expression, order []clause.OrderByColumn, opts ...OptionFunc,
) (*Page[T], *gorm.DB) {
	return r.PageByClause(where, order, nil, opts...)
}

func (r *Repository[T]) Paginate(page *Page[T], opts ...OptionFunc) (*Page[T], *gorm.DB) {
	entities := make([]*T, 0)
	db := r.GetDB(opts...)

	if page != nil {
		if !page.ignoreTotal {
			total, cdb := r.Count(opts...)
			if cdb.Error != nil {
				return page, cdb
			}

			page.total = total
		}

		db = db.Scopes(page.Paginate())
	} else {
		page = NewPage[T](1, 0)
	}

	db = db.Model(r.model).Find(&entities)

	page.Records = entities

	return page, db
}

func (r *Repository[T]) PageByCondition(
	cond func(*gorm.DB) *gorm.DB, page *Page[T], opts ...OptionFunc,
) (*Page[T], *gorm.DB) {
	entities := make([]*T, 0)
	db := r.GetDB(opts...)

	if page != nil {
		if !page.ignoreTotal {
			total, cdb := r.CountByCondition(cond, opts...)
			if cdb.Error != nil {
				return page, cdb
			}

			page.total = total
		}

		db = db.Scopes(page.Paginate())
	} else {
		page = NewPage[T](1, 0)
	}

	db = db.Scopes(cond).Find(&entities)

	page.Records = entities

	return page, db
}

func (r *Repository[T]) PageByClause(
	where []clause.Expression, order []clause.OrderByColumn, page *Page[T], opts ...OptionFunc,
) (*Page[T], *gorm.DB) {
	entities := make([]*T, 0)
	db := r.GetDB(opts...)

	db = db.Clauses(clause.Where{Exprs: where}, clause.OrderBy{Columns: order})

	if page != nil {
		if !page.ignoreTotal {
			total, cdb := r.CountByClause(where, opts...)
			if cdb.Error != nil {
				return page, cdb
			}

			page.total = total
		}

		db = db.Scopes(page.Paginate())
	} else {
		page = NewPage[T](1, 0)
	}

	db.Find(&entities)

	page.Records = entities

	return page, db
}

func (r *Repository[T]) CursorPageByCondition(
	cond func(*gorm.DB) *gorm.DB, page *CursorPage[T], opts ...OptionFunc,
) (*CursorPage[T], *gorm.DB) {
	entities := make([]*T, 0)
	db := r.GetDB(opts...)

	if !page.ignoreTotal {
		total, cdb := r.CountByCondition(cond, opts...)
		if cdb.Error != nil {
			return page, cdb
		}

		page.total = total
	}

	db = db.Scopes(page.Paginate(), cond).Find(&entities)

	page.Records = entities

	return page, db
}

func (r *Repository[T]) Count(opts ...OptionFunc) (int64, *gorm.DB) {
	var (
		db  = r.GetDB(opts...)
		cnt int64
	)

	db = db.Model(r.model).Count(&cnt)

	return cnt, db
}

func (r *Repository[T]) CountByCondition(cond func(*gorm.DB) *gorm.DB, opts ...OptionFunc) (int64, *gorm.DB) {
	var (
		db  = r.GetDB(opts...)
		cnt int64
	)

	db = db.Model(r.model).Scopes(cond).Count(&cnt)

	return cnt, db
}

func (r *Repository[T]) CountByClause(where []clause.Expression, opts ...OptionFunc) (int64, *gorm.DB) {
	var (
		db  = r.GetDB(opts...)
		cnt int64
	)

	db = db.Model(r.model).Clauses(clause.Where{Exprs: where}).Count(&cnt)

	return cnt, db
}

// Pluck queries a single column from a model, returning in the slice dest.
func (r *Repository[T]) Pluck(col string, dest any, opts ...OptionFunc) *gorm.DB {
	db := r.GetDB(opts...)

	return db.Pluck(col, dest)
}

func (r *Repository[T]) GetDB(opts ...OptionFunc) *gorm.DB {
	opt := getOption(opts)

	db := r.db

	if opt.tx != nil {
		db = opt.tx
	}

	if len(opt.selects) > 0 {
		db = db.Select(opt.selects)
	}

	if len(opt.omits) > 0 {
		db = db.Omit(opt.omits...)
	}

	if opt.order != "" {
		db = db.Order(opt.order)
	}

	return db
}

// NewSession create new db session.
func (r *Repository[T]) NewSession() *gorm.DB {
	return r.db.Session(&gorm.Session{})
}

func (r *Repository[T]) TableName() string {
	getTableName := func(modelValue reflect.Value) string {
		tableName := ""
		if t, ok := modelValue.Interface().(schema.Tabler); ok {
			tableName = t.TableName()
		}

		if t, ok := modelValue.Interface().(schema.TablerWithNamer); ok {
			tableName = t.TableName(schema.NamingStrategy{IdentifierMaxLength: 64})
		}

		return tableName
	}

	modelValue := reflect.ValueOf(r.model)
	tableName := getTableName(modelValue)

	if tableName == "" {
		modelValue = reflect.Indirect(modelValue)
		tableName = getTableName(modelValue)

		if tableName == "" {
			tableName = modelValue.Type().Name()
		}
	}

	return tableName
}
