package orm

import (
	"fmt"

	"gorm.io/gorm"
)

func IsEmpty(v any) bool {
	switch v.(type) {
	case int:
		return v == 0
	case int8:
		return v == int8(0)
	case int16:
		return v == int16(0)
	case int32:
		return v == int32(0)
	case int64:
		return v == int64(0)
	case uint:
		return v == uint(0)
	case uint8:
		return v == uint8(0)
	case uint16:
		return v == uint16(0)
	case uint32:
		return v == uint32(0)
	case uint64:
		return v == uint64(0)
	case string:
		return v == ""
	}

	return false
}

type SortOrder string

const (
	OrderAsc  SortOrder = "ASC"
	OrderDesc SortOrder = "DESC"
)

func (c SortOrder) RawValue() string {
	return string(c)
}

type Page[T any] struct {
	page        int
	limit       int
	total       int64
	ignoreTotal bool
	order       string

	Records    []*T
	RecordsMap []T
}

func NewPage[T any](page, limit int64) *Page[T] {
	return &Page[T]{
		page:       int(page),
		limit:      int(limit),
		Records:    make([]*T, 0),
		RecordsMap: make([]T, 0),
	}
}

func NewPageNoTotal[T any](page, limit int64) *Page[T] {
	return &Page[T]{
		page:        int(page),
		limit:       int(limit),
		ignoreTotal: true,
		Records:     make([]*T, 0),
		RecordsMap:  make([]T, 0),
	}
}

func (p *Page[T]) Total() int64 {
	return p.total
}

func (p *Page[T]) SetOrder(order string) {
	p.order = order
}

func (p *Page[T]) Paginate() func(db *gorm.DB) *gorm.DB {
	page := p.page
	pageSize := p.limit

	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		if pageSize <= 0 {
			pageSize = 10
		}

		offset := (page - 1) * pageSize

		return db.Offset(offset).Limit(pageSize)
	}
}

type CursorPage[T any] struct {
	column      string
	sort        SortOrder
	start       any
	limit       int
	total       int64
	ignoreTotal bool

	Records    []*T `json:"records"`
	RecordsMap []T  `json:"recordsMap"`
}

func NewCursorPage[T any](col string, s SortOrder, start any, limit int64) *CursorPage[T] {
	return &CursorPage[T]{
		column:     col,
		start:      start,
		sort:       s,
		limit:      int(limit),
		Records:    make([]*T, 0),
		RecordsMap: make([]T, 0),
	}
}

func (p *CursorPage[T]) Total() int64 {
	return p.total
}

func (p *CursorPage[T]) Paginate() func(db *gorm.DB) *gorm.DB {
	column := p.column
	start := p.start
	pageSize := p.limit
	st := p.sort

	return func(db *gorm.DB) *gorm.DB {
		if pageSize <= 0 {
			pageSize = 10
		}

		db = db.Order(fmt.Sprintf("%s %s", column, st.RawValue())).
			Limit(pageSize)

		if !IsEmpty(start) {
			op := ">"
			if st == OrderDesc {
				op = "<"
			}

			db = db.Where(fmt.Sprintf("%s %s ?", column, op), start)
		}

		return db
	}
}
