package tp4_calculdora

import (
	"fmt"

	"github.com/rodri/Lab3-master/Lab3/tp4_calculadora/operaciones"
)

func Calculadora() {
	var num1, num2 float64
	var operacion int
	var operandos []float64
	var respuesta string
	var bug bool
	var offset int
	fmt.Printf("Ingrese la operación a realizar: \n")
	fmt.Printf("1 : Sumar \n")
	fmt.Printf("2 : Restar \n")
	fmt.Printf("3 : Multipl. \n")
	fmt.Printf("4 : Dividir \n")
	fmt.Scan(&operacion)
	fmt.Printf(" - Ingrese dos operandos: \n")
	fmt.Scan(&num1, &num2)
	operandos = []float64{num1, num2}
	fmt.Printf("Desea buguear la calculadora?: (S/N)")
	fmt.Scan(&respuesta)
	if respuesta == "S" {
		bug = true
		fmt.Printf("Ingrese OFFSET del Bug:")
		fmt.Scan(&offset)
	}
	switch operacion {
	case 1:
		fmt.Print(operaciones.Sumar(bug, offset, operandos...))
	case 2:
		fmt.Print(operaciones.Restar(bug, offset, operandos...))
	case 3:
		fmt.Print(operaciones.Multiplicar(bug, offset, operandos...))
	case 4:
		fmt.Print(operaciones.Division(bug, offset, operandos...))
	default:
		fmt.Print("Opcion no soportada")
	}
}
