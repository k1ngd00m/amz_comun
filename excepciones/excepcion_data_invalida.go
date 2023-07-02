package excepciones

const DATA_INVALIDA = "datos invalidos"

func DataInvalida(mensaje string) error {
	return &ErrorCustom{
		mensaje:     mensaje,
		categoria:   DATA_INVALIDA,
		codigo:      400,
		operacional: false,
	}
}
