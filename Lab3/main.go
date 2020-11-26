package main

import (
	"bufio"
	"fmt"
	"os"

	tp3_hello_word "github.com/rodri/Lab3-master/Lab3/tp3_holamundo"
	tp4_calculdora "github.com/rodri/Lab3-master/Lab3/tp4_calculadora"
	tp5_consumo_apis "github.com/rodri/Lab3-master/Lab3/tp5_Apis"
	tp6_calculadora_hilos "github.com/rodri/Lab3-master/Lab3/tp6_CalculadoraHilos"
	tp7_calculadora_test "github.com/rodri/Lab3-master/Lab3/tp7Calculadora_test"
	tp8_calculadoraApi "github.com/rodri/Lab3-master/Lab3/tp8_calculadoraEnApi"
)

func main() {
	var numeroTP int64
	var numeroTP2 bool
	//Se muestra el menu de opciones con los TP incluidos en el proyecto en forma de pkg

	for numeroTP2 = true; numeroTP2; {
		fmt.Printf("\n\n ")
		fmt.Printf("*********************************\n")
		fmt.Printf(" 3 : HolaMundo \n")
		fmt.Printf(" 4 : Variables y estructuras en Go \n")
		fmt.Printf(" 5 : Consumo de APIs en Go \n")
		fmt.Printf(" 6 : Calculadora // Concurrencia y Paralelirmos en Go \n")
		fmt.Printf(" 7 : Testing \n")
		fmt.Printf(" 8 : Creacion de una API \n\n ")
		fmt.Printf(" 9 : Cerrar Programa\n\n ")
		fmt.Printf(" - Elija TP que quiere ejecutar - \n")
		fmt.Printf("*********************************\n")

		fmt.Scan(&numeroTP)
		switch numeroTP {
		case 3:
			fmt.Printf("\n\n ")
			fmt.Printf("========= - TP 3 - ========\n")
			fmt.Printf("\n\n ")
			tp3_hello_word.HolaMundo()
			fmt.Printf("\n\n ")
			fmt.Print("Presione 'Enter' para continuar...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			break
		case 4:
			fmt.Printf("\n\n ")
			fmt.Printf("========= - TP 4 - ========\n")
			fmt.Printf("\n\n ")
			tp4_calculdora.Calculadora()
			fmt.Printf("\n\n ")
			fmt.Print("Presione 'Enter' para continuar...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			break
		case 5:
			fmt.Printf("\n\n ")
			fmt.Printf("========= - TP 5 - ========\n")
			fmt.Printf("\n\n ")
			tp5_consumo_apis.ConsumoApi()
			fmt.Printf("\n\n ")
			fmt.Print("Presione 'Enter' para continuar...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			break
		case 6:
			fmt.Printf("\n\n ")
			fmt.Printf("========= - TP 6 - ========\n")
			fmt.Printf("\n\n ")
			tp6_calculadora_hilos.CalculadoraHilos()
			fmt.Printf("\n\n ")
			fmt.Print("Presione 'Enter' para continuar...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')

			break
		case 7:
			fmt.Printf("\n\n ")
			fmt.Printf("========= - TP 7 - ========\n")
			fmt.Printf("\n\n ")
			tp7_calculadora_test.CalculadoraTest()
			fmt.Printf("\n\n ")
			fmt.Print("Presione 'Enter' para continuar...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			break
		case 8:
			fmt.Printf("\n\n ")
			fmt.Printf("========= - TP 8 - ========\n")
			fmt.Printf("\n\n ")
			tp8_calculadoraApi.CalculadoraApi()
			fmt.Printf("\n\n ")
			fmt.Print("Presione 'Enter' para continuar...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			break
		case 9:
			fmt.Printf(" - Programa Finalizado - \n")
			numeroTP2 = false
		default:
			fmt.Printf(" - Opcion no soportada, vuelva a ingresar - \n")
		}
	}
}
