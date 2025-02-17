package orm

import "gorm.io/gorm"

type Option struct {
	selects []string
	omits   []string
	order   string
	tx      *gorm.DB
}

type OptionFunc func(*Option)

func getOption(opts []OptionFunc) Option {
	var opt Option

	for _, op := range opts {
		op(&opt)
	}

	return opt
}

func Select(cols ...string) OptionFunc {
	return func(o *Option) {
		o.selects = append(o.selects, cols...)
	}
}

func Omit(cols ...string) OptionFunc {
	return func(o *Option) {
		o.omits = append(o.omits, cols...)
	}
}

func Order(order string) OptionFunc {
	return func(o *Option) {
		o.order = order
	}
}

func TX(tx *gorm.DB) OptionFunc {
	return func(o *Option) {
		o.tx = tx
	}
}
