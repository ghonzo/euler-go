// Project Euler
package common

// Returns true if there are any repeated elements
func HasRepeatedElements[T comparable](slice []T) bool {
	seen := make(map[T]bool)
	for _, v := range slice {
		if seen[v] {
			return true
		}
		seen[v] = true
	}
	return false
}