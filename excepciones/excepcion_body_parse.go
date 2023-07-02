package excepciones

const dataParsing = "parsing body"

const parseoJson = "parseo json"

func ParseoBody(mensaje string) error {
	return &ErrorCustom{
		mensaje:     mensaje,
		categoria:   parseoJson,
		codigo:      400,
		operacional: false,
	}
}
