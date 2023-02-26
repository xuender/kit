package logs

import (
	"io"
	"log"
	"os"
)

type logger struct {
	logger *log.Logger
	output io.Writer
}

func (p *logger) newLog(prefix string) *log.Logger {
	p.logger = log.New(os.Stderr, prefix, log.Ltime|log.Lshortfile)

	return p.logger
}

func (p *logger) setOutput(writer io.Writer) {
	p.output = writer
	p.logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func (p *logger) reset() {
	p.logger.SetOutput(p.output)
}

func (p *logger) ignore() {
	p.logger.SetOutput(io.Discard)
}
