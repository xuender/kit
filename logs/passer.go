package logs

type passer struct{}

func (p *passer) Write(data []byte) (int, error) {
	return len(data), nil
}
