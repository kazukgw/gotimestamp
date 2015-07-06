package gotimestamp

import (
	"time"
)

var location *time.Location

func SetLocation(loc *time.Location) {
	location = loc
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
	t.CreatedAtInLocal = t.CreatedAt.In(location).Format(format)
	t.UpdatedAtInLocal = t.UpdatedAt.In(location).Format(format)
}

func (t *TimeStamp) ParseInLocation(createdat, updatedat string) error {
	var err error
	if createdat != "" {
		t.CreatedAt, err = time.ParseInLocation(format, createdat, location)
	}
	if updatedat != "" {
		t.UpdatedAt, err = time.ParseInLocation(format, updatedat, location)
	}
	return err
}
