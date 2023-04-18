package ustring

import (
	"fmt"
	"time"
)

// ToDate converte uma string para data (dd/mm/yyyy).
func ToDate(strDate string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", strDate)
	if err != nil {
		return date, fmt.Errorf("não foi possível converter o texto para data")
	}

	return date, nil
}
