package umap

// Copy retorna a cópia de um map.
func Copy[T any](m map[string]T) map[string]T {
	newMap := make(map[string]T)

	for k, v := range m {
		newMap[k] = v
	}

	return newMap
}
