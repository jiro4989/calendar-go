package calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsSameDay(t *testing.T) {
	tests := []struct {
		desc string
		a, b time.Time
		want bool
	}{
		{
			desc: "ok: 2021/12/01 00:00:00 equals 2021/12/01 15:00:00",
			a:    time.Date(2021, time.December, 1, 0, 0, 0, 0, time.Local),
			b:    time.Date(2021, time.December, 1, 15, 0, 0, 0, time.Local),
			want: true,
		},
		{
			desc: "ok: 2021/12/01 00:00:00 equals 2021/12/02 00:00:00",
			a:    time.Date(2021, time.December, 1, 0, 0, 0, 0, time.Local),
			b:    time.Date(2021, time.December, 2, 0, 0, 0, 0, time.Local),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			got := IsSameDay(tt.a, tt.b)
			assert.Equal(tt.want, got)
		})
	}
}

func TestContains(t *testing.T) {
	day1 := time.Date(2021, time.December, 1, 0, 0, 0, 0, time.Local)

	tests := []struct {
		desc    string
		day     time.Time
		days    []time.Time
		want    bool
		wantDay *time.Time
	}{
		{
			desc: "ok: true",
			day:  day1,
			days: []time.Time{
				time.Date(2021, time.December, 1, 0, 0, 0, 0, time.Local),
				time.Date(2021, time.December, 2, 0, 0, 0, 0, time.Local),
				time.Date(2021, time.December, 3, 0, 0, 0, 0, time.Local),
			},
			want:    true,
			wantDay: &day1,
		},
		{
			desc: "ok: false",
			day:  day1,
			days: []time.Time{
				time.Date(2021, time.December, 2, 0, 0, 0, 0, time.Local),
				time.Date(2021, time.December, 3, 0, 0, 0, 0, time.Local),
			},
			want:    false,
			wantDay: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			found, got := Contains(tt.day, tt.days)
			assert.Equal(tt.want, found)
			assert.Equal(tt.wantDay, got)
		})
	}
}
