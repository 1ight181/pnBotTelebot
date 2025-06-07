package sliceutils

func RemoveByValue[T comparable](slice []T, valueToRemove T) []T {
	for i, value := range slice {
		if value == valueToRemove {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
