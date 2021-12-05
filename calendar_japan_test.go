package calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func jd(y int, m time.Month, d int, r Rokuyo) JapanDay {
	return JapanDay{
		Day:    time.Date(y, m, d, 0, 0, 0, 0, time.Local),
		Rokuyo: r,
	}
}

func TestCalendarDaysOfJapan(t *testing.T) {
	tests := []struct {
		desc  string
		year  int
		month time.Month
		want  []JapanDay
	}{
		{
			desc:  "ok: 2021/12/1",
			year:  2021,
			month: time.December,
			want: []JapanDay{
				jd(2021, time.November, 28, Tomobiki),
				jd(2021, time.November, 29, Senpu),
				jd(2021, time.November, 30, Butsumetsu),
				jd(2021, time.December, 1, Shakko),
				jd(2021, time.December, 2, Sensho),
				jd(2021, time.December, 3, Tomobiki),
				jd(2021, time.December, 4, Senpu),
				jd(2021, time.December, 5, Butsumetsu),
				jd(2021, time.December, 6, Taian),
				jd(2021, time.December, 7, Shakko),
				jd(2021, time.December, 8, Sensho),
				jd(2021, time.December, 9, Tomobiki),
				jd(2021, time.December, 10, Senpu),
				jd(2021, time.December, 11, Butsumetsu),
				jd(2021, time.December, 12, Taian),
				jd(2021, time.December, 13, Shakko),
				jd(2021, time.December, 14, Sensho),
				jd(2021, time.December, 15, Tomobiki),
				jd(2021, time.December, 16, Senpu),
				jd(2021, time.December, 17, Butsumetsu),
				jd(2021, time.December, 18, Taian),
				jd(2021, time.December, 19, Shakko),
				jd(2021, time.December, 20, Sensho),
				jd(2021, time.December, 21, Tomobiki),
				jd(2021, time.December, 22, Senpu),
				jd(2021, time.December, 23, Butsumetsu),
				jd(2021, time.December, 24, Taian),
				jd(2021, time.December, 25, Shakko),
				jd(2021, time.December, 26, Sensho),
				jd(2021, time.December, 27, Tomobiki),
				jd(2021, time.December, 28, Senpu),
				jd(2021, time.December, 29, Butsumetsu),
				jd(2021, time.December, 30, Taian),
				jd(2021, time.December, 31, Shakko),
				jd(2022, time.January, 1, Sensho),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			got := CalendarDaysOfJapan(tt.year, tt.month)
			assert.Equal(tt.want, got)
		})
	}
}

func TestToJapanDay(t *testing.T) {
	tests := []struct {
		desc string
		t    time.Time
		want JapanDay
	}{
		{
			desc: "ok: 10月1日は仏滅",
			t:    time.Date(2021, 10, 1, 0, 0, 0, 0, time.Local),
			want: jd(2021, 10, 1, Butsumetsu),
		},
		{
			desc: "ok: 11月1日は大安",
			t:    time.Date(2021, 11, 1, 0, 0, 0, 0, time.Local),
			want: jd(2021, 11, 1, Taian),
		},
		{
			desc: "ok: 11月2日は赤口",
			t:    time.Date(2021, 11, 2, 0, 0, 0, 0, time.Local),
			want: jd(2021, 11, 2, Shakko),
		},
		{
			desc: "ok: 11月3日は先勝",
			t:    time.Date(2021, 11, 3, 0, 0, 0, 0, time.Local),
			want: jd(2021, 11, 3, Sensho),
		},
		{
			desc: "ok: 12月1日は赤口",
			t:    time.Date(2021, 12, 1, 0, 0, 0, 0, time.Local),
			want: jd(2021, 12, 1, Shakko),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			got := ToJapanDay(tt.t)
			assert.Equal(tt.want, got)
		})
	}
}

func TestRokuyo_Next(t *testing.T) {
	tests := []struct {
		desc string
		r    Rokuyo
		want Rokuyo
	}{
		{
			desc: "ok: 先勝の次は友引",
			r:    Sensho,
			want: Tomobiki,
		},
		{
			desc: "ok: 赤口の次は先勝",
			r:    Shakko,
			want: Sensho,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			got := tt.r.next()
			assert.Equal(tt.want, got)
		})
	}
}

func TestRokuyo_Kanji(t *testing.T) {
	tests := []struct {
		desc string
		r    Rokuyo
		want string
	}{
		{
			desc: "ok: Sensho is 先勝",
			r:    Sensho,
			want: "先勝",
		},
		{
			desc: "ok: Tomobiki is 友引",
			r:    Tomobiki,
			want: "友引",
		},
		{
			desc: "ok: Senpu is 先負",
			r:    Senpu,
			want: "先負",
		},
		{
			desc: "ok: Butsumetsu is 仏滅",
			r:    Butsumetsu,
			want: "仏滅",
		},
		{
			desc: "ok: Taian is 大安",
			r:    Taian,
			want: "大安",
		},
		{
			desc: "ok: Shakko is 赤口",
			r:    Shakko,
			want: "赤口",
		},
		{
			desc: "ok: 6 is 不明",
			r:    6,
			want: "不明",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			got := tt.r.Kanji()
			assert.Equal(tt.want, got)
		})
	}
}
