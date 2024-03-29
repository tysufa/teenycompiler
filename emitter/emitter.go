package emitter

import (
	"os"
)

type Emitter struct {
	fullPath string
	header   string
	code     string
}

func New(fullPath string) Emitter {
	return Emitter{fullPath: fullPath}
}

func (e *Emitter) Emit(code string) {
	e.code += code
}

func (e *Emitter) EmitLine(code string) {
	e.code += code + "\n"
}

func (e *Emitter) HeaderLine(code string) {
	e.header += code + "\n"
}

func (e *Emitter) WriteFile() {
	f, err := os.Create(e.fullPath)
	if err != nil {
		panic(err)
	}

	f.WriteString(e.header + e.code)
	f.Close()
}
