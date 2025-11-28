// Package cache provides utilities for caching data.
package cache

// Memoize is a generic function that takes a key, a cache, and a function.
// If the key is in the cache, the value is returned. Otherwise, the function
// is called with the key, the result is stored in the cache, and the result
// is returned.
func Memoize[T comparable, U any](key T, cache map[T]U, fn func(T) U) U {
	if val, ok := cache[key]; ok {
		return val
	}

	val := fn(key)
	cache[key] = val

	return val
}
