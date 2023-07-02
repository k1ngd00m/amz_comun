package excepciones

type ErrorCustom struct {
	codigo      int
	categoria   string
	mensaje     string
	operacional bool
}

func (e *ErrorCustom) Codigo() int {
	return e.codigo
}

func (e *ErrorCustom) Operacional() int {
	return e.codigo
}

func (e *ErrorCustom) Categoria() string {
	return e.categoria
}

func (e *ErrorCustom) Error() string {
	return e.mensaje
}
