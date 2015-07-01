package gotimestamp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateTimeStamp(t *testing.T) {
	a := assert.New(t)
	ts := &TimeStamp{}
	ts.CreateTimeStamp()
	a.NotEqual(time.Time{}, ts.CreatedAt, "ts.CreatedAt != zero after call 'CreateTimeStamp'")
	a.Equal(ts.CreatedAt.Unix(), ts.UpdatedAt.Unix(), "created_at == updated_at after call 'CreateTimeStamp'")
}

func TestComputeLocalTime(t *testing.T) {
	a := assert.New(t)
	ts := &TimeStamp{}
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		loc = time.FixedZone("Asia/Tokyo", 9*60*60)
	}
	_t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	ts.CreatedAt = _t
	ts.UpdatedAt = _t
	SetFormat("2006-01-02 15:04:05")
	SetLocation(loc)
	ts.ComputeLocalTime()
	a.Equal("2009-11-11 08:00:00", ts.CreatedAtInLocal, "CreatedAtInLocal after call 'ComputeLocalTime'")
}
