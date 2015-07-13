package gotimestamp

import (
	"time"
)

var Location *time.Location

func SetLocation(loc *time.Location) {
	Location = loc
}

var Format string

func SetFormat(f string) {
	Format = f
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
	t.CreatedAtInLocal = t.CreatedAt.In(Location).Format(Format)
	t.UpdatedAtInLocal = t.UpdatedAt.In(Location).Format(Format)
}

func (t *TimeStamp) ComputeLocalTimeWithLocation(l *time.Location) {
	t.CreatedAtInLocal = t.CreatedAt.In(l).Format(Format)
	t.UpdatedAtInLocal = t.UpdatedAt.In(l).Format(Format)
}

func (t *TimeStamp) ParseInLocation(createdat, updatedat string) error {
	var err error
	if createdat != "" {
		t.CreatedAt, err = time.ParseInLocation(Format, createdat, Location)
	}
	if updatedat != "" {
		t.UpdatedAt, err = time.ParseInLocation(Format, updatedat, Location)
	}
	return err
}
