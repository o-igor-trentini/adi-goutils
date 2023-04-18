package uenv

import (
	"fmt"
	"github.com/o-igor-trentini/adi-goutils/pkg/uslice"
	"os"
	"strings"
)

// GetStr retorna o valor da variável de ambiente como string ou erro caso o valor esteja vazio.
func GetStr(name string) (string, error) {
	v := os.Getenv(name)
	if v == "" {
		return "", fmt.Errorf("valor não encontrado para a env: %s", name)
	}

	return v, nil
}

// CheckEnvs verifica se os valores de todas as variáveis de ambiente foram informadadas.
func CheckEnvs(table []CheckEnvTable) error {
	var (
		missingEnvs []string
		wrongValues []string
	)
	for _, item := range table {
		v, err := GetStr(item.Name)
		if err != nil && (v == "" && !item.IgnoreEmpty) {
			missingEnvs = append(missingEnvs, item.Name)
		}

		if len(item.Expected) > 0 {
			if !uslice.Contains(item.Expected, v) {
				wrongValues = append(wrongValues, fmt.Sprintf("%s=%s", item.Name, v))
			}
		}
	}

	if len(missingEnvs) > 0 {
		return fmt.Errorf(
			"não foi possível carregar as seguintes variáveis de ambiente: [%s]",
			strings.Join(missingEnvs, ", "),
		)
	}

	if len(wrongValues) > 0 {
		return fmt.Errorf(
			"as seguintes variáveis de ambiente não receberam os valores esperados: [%s]",
			strings.Join(wrongValues, ", "),
		)
	}

	return nil
}
