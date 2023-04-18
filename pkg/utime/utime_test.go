package utime_test

import (
	"github.com/o-igor-trentini/adi-goutils/pkg/utime"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRemoveTime(t *testing.T) {
	name := "Remover o horário com sucesso"
	actual := time.Date(2023, 2, 15, 12, 30, 45, 99, time.Local)
	want := time.Date(2023, 2, 15, 0, 0, 0, 0, time.Local)

	if want != utime.RemoveTime(actual) {
		assert.Equal(t, want, actual, name)
	}
}
