package today

import "time"

func String(t time.Time) string {
	return t.Format("2006-01-02")
}
