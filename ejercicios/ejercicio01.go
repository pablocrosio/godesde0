package ejercicios

import "strconv"

func Convertir(texto string) (int, string) {
	numero, error := strconv.Atoi(texto)
	if error != nil {
		return 0, "Hubo un error" + error.Error()
	}
	if numero > 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "Es menor o igual a 100"
	}
}
