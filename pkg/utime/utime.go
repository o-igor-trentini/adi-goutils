package utime

import (
	"time"
)

// RemoveTime remove as horas, mantendo apenas a data.
func RemoveTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
