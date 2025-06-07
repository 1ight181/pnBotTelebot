package sliceutils

func In[T comparable](val T, slice []T) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
