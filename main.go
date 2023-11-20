package main

import (
	"fmt"
	//"runtime"
	//"github.com/pablocrosio/godesde0/variables"
	"github.com/pablocrosio/godesde0/ejercicios"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)*/
	/*if os := runtime.GOOS; os == "linux" {
		fmt.Println("Esto no es Windows, es ", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}*/
	numero, mensaje := ejercicios.Convertir("455")
	fmt.Println("numero = ", numero)
	fmt.Println("mensaje = ", mensaje)
}
