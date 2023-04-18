package uslice_test

import (
	"github.com/o-igor-trentini/adi-goutils/pkg/uslice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	type testCase[T any] struct {
		name         string
		inputSlice   []T
		inputElement T
		want         bool
	}

	strTest := []testCase[string]{
		{
			name:         "Slice de string contêm valor",
			inputSlice:   []string{"firstElement", "secondElement", "thirdElement"},
			inputElement: "secondElement",
			want:         true,
		},
		{
			name:         "slice de string não contêm valor",
			inputSlice:   []string{"firstElement", "secondElement", "thirdElement"},
			inputElement: "Element",
			want:         false,
		},
	}

	intTest := []testCase[int]{
		{
			name:         "Slice de int contêm valor",
			inputSlice:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			inputElement: 4,
			want:         true,
		},
		{
			name:         "Slice de int não contêm valor",
			inputSlice:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			inputElement: 11,
			want:         false,
		},
	}

	// Manter mesma quantidade de testes
	for i := range strTest {
		assert.Equal(
			t,
			strTest[i].want,
			uslice.Contains(strTest[i].inputSlice, strTest[i].inputElement),
			strTest[i].name,
		)

		assert.Equal(
			t,
			intTest[i].want,
			uslice.Contains(intTest[i].inputSlice, intTest[i].inputElement),
			intTest[i].name,
		)
	}
}
