/*
日本のカレンダー関連。
六曜は旧暦が基準になっているため、旧暦←→グレゴリオ暦の変換が必要になる。

参考: https://www.ndl.go.jp/koyomi/chapter3/s3.html

*/
package calendar

import "time"

type Rokuyo int // 六曜

type JapanDay struct {
	Day    time.Time
	Rokuyo Rokuyo
}

const (
	Sensho     Rokuyo = iota // 先勝
	Tomobiki                 // 友引
	Senpu                    // 先負
	Butsumetsu               // 仏滅
	Taian                    // 大安
	Shakko                   // 赤口
)

var (
	// 先勝→友引→先負→仏滅→大安→赤口の順で繰り返すが、毎月1日は月によって決
	// まっている。毎月1日の方が優先度が高い。
	startRokuyo = map[time.Month]Rokuyo{
		time.January:   Sensho,     // 1月 先勝
		time.February:  Tomobiki,   // 2月 友引
		time.March:     Senpu,      // 3月 先負
		time.April:     Butsumetsu, // 4月 仏滅
		time.May:       Taian,      // 5月 大安
		time.June:      Shakko,     // 6月 赤口
		time.July:      Sensho,     // 7月 先勝
		time.August:    Tomobiki,   // 8月 友引
		time.September: Senpu,      // 9月 先負
		time.October:   Butsumetsu, // 10月 仏滅
		time.November:  Taian,      // 11月 大安
		time.December:  Shakko,     // 12月 赤口
	}
)

// CalendarDaysOfJapan returns days like `cal` calendar for japan.
func CalendarDaysOfJapan(y int, m time.Month) []JapanDay {
	var result []JapanDay
	days := CalendarDays(y, m)
	for i := range days {
		day := days[i]
		jd := ToJapanDay(day)
		result = append(result, jd)
	}
	return result
}

// ToJapanDay returns a Japanese day.
func ToJapanDay(t time.Time) JapanDay {
	sr := startRokuyo[t.Month()]
	for i := 1; i < t.Day(); i++ {
		sr = sr.next()
	}
	jd := JapanDay{
		Day:    t,
		Rokuyo: sr,
	}
	return jd
}

// next は r の次の六曜を返す。
func (r Rokuyo) next() Rokuyo {
	switch r {
	case Shakko:
		r = Sensho
	default:
		r++
	}
	return r
}

// Kanji returns a japanese kanji (漢字).
func (r Rokuyo) Kanji() string {
	switch r {
	case Sensho:
		return "先勝"
	case Tomobiki:
		return "友引"
	case Senpu:
		return "先負"
	case Butsumetsu:
		return "仏滅"
	case Taian:
		return "大安"
	case Shakko:
		return "赤口"
	}
	return "不明" // 本来であれば到達しない
}
