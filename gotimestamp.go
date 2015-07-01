package gotimestamp

import (
	"time"
)

var local time.Location

func SetLocal(loc time.Location) {
	local = loc
}

var format string

func SetFormat(f string) {
	format = f
}

type TimeStamp struct {
	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`

	CreatedAtInLocal string `db:"-" json:"created_at"`
	UpdatedAtInLocal string `db:"-" json:"updated_at"`
}

func (t *TimeStamp) CreateTimeStamp() {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}

func (t *TimeStamp) UpdateTimeStamp() {
	t.UpdatedAt = time.Now()
}

func (t *TimeStamp) ComputeLocalTime() {
	t.CreatedAtInLocal = t.CreatedAt.In(local).format(format)
	t.UpdatedAtInLocal = t.UpdatedAt.In(local).format(format)
}
