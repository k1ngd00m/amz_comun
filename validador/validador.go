package validadores

import "github.com/k1ngd00m/amz_comun/excepciones"

func StringObligatorio(cadena string, mensaje string) error {
	if len(cadena) == 0 {
		return excepciones.DataInvalida(mensaje)
	}

	return nil
}

func NumeroMayorCero(numero int16, mensaje string) error {
	if numero < 0 {
		return excepciones.DataInvalida(mensaje)
	}
	return nil
}
