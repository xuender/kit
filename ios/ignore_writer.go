package ios

type IgnoreWriter struct{}

func (p IgnoreWriter) Write(data []byte) (int, error) {
	return len(data), nil
}
