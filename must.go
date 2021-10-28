package must

// Must provides a convenient way to do a "must" pattern in Go in generic way.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
