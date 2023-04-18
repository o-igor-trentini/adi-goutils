package uslice

// Contains verifica se o valor existe no slice.
func Contains[T comparable](slice []T, element T) bool {
	for _, s := range slice {
		if element == s {
			return true
		}
	}

	return false
}

// Index busca o indíce do valor no slice com base no valor passado, caso não exista é retornado '-1'.
func Index[T comparable](slice []T, element T) int {
	for k, s := range slice {
		if element == s {
			return k
		}
	}

	return -1
}
