package emitter

type Emitter struct {
	fullPath string
	header   string
	code     string
}

func New(fullPath string) Emitter {
	return Emitter{fullPath: fullPath}
}

func (e *Emitter) emit(code string) {
	e.code += code + "\n"
}
