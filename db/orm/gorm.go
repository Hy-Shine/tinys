package orm

import (
	"gorm.io/gorm"
)

type DBFunc func(*gorm.DB)

func WithLRLike(column, value string) DBFunc {
	return func(d *gorm.DB) {
		if column != "" && value != "" {
			d.Where(column+" LIKE ?", "%"+value+"%")
		}
	}
}

func WithLike(column, value string) DBFunc {
	return func(d *gorm.DB) {
		if column != "" && value != "" {
			d.Where(column+" LIKE ?", value)
		}
	}
}

func WithWhere(exp string, args ...any) DBFunc {
	return func(d *gorm.DB) {
		d.Where(exp, args...)
	}
}

func WithBetween[T comparable](column string, begin, end T) DBFunc {
	return func(d *gorm.DB) {
		d.Where(column+" ? AND ?", begin, end)
	}
}

func WithValue(column string, value any) DBFunc {
	return func(d *gorm.DB) {
		if value != nil {
			d.Where(column+" = ?", value)
		}
	}
}

func WithValueIn(column string, value any) DBFunc {
	return func(d *gorm.DB) {
		if value != nil {
			d.Where(column+" IN (?)", value)
		}
	}
}
