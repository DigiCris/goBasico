package main

import "fmt"
import "bufio"
import "os"


/* Ejercicio

1) Impresión de textos y variables por consola:
	a) Al solicitar al usuario que ingrese el radio de un circulo.
	b) Al imprimir los resultados del área y perímetro.
2) Uso de diferentes tipos de variables y constantes:
	a) Variables enteras para el radio.
	b) Variables flotantes para calcular y almacenar el área y perímetro.
	c) Constante PI flotante para utilizar en el cálculo del perímetro.
3) Función que recibe argumentos y devuelve más de un valor:
	a) La función que calcula el área y perímetro.
4) Uso de al menos un bucle for:
	a) validar que los radios ingresados sean positivos y enteros
	b) Validar que estamos ingresando opciones validas de impresion
5) Uso de al menos un operador || o &&:
	a) Se utilizará el operador && para validar que los valores ingresados sean enteros positivos.
6) Uso de al menos un if:
	a) Se utilizará un if para distinguir el error de ingreso de radio
7) Uso de al menos un switch:
	a) Se utilizará un switch para decidir qué resultado imprimir (área o perímetro) después de los cálculos.
8) Uso de defer:
	a) Un mensaje de fin para cuando termina de ejecutar el main que se pondrá al principio

*/

func calculaAreaPerimetro(radio uint) (area, perimetro float32){ //3)
	const pi float32 = 3.1415 // 2)
	area = pi * float32(radio) * float32(radio) // 2)
	perimetro = pi * (2*float32(radio)) // 2)
	return area, perimetro 
}

func main(){
	var radio uint //2)
	defer fmt.Println("Termine de calcular todo"); // 8)
	fmt.Println("Ingrese el valor del radio del circulo: ");
	_, err := fmt.Scanln(&radio); // si err == nil entonces está bien

	for i:=0; err !=nil || radio<0 ; i++ { // 4) y 5)
		if err.Error()=="expected integer" { // 6)
			fmt.Println("tiene que ser un numero entero positivo")
		}else {
			fmt.Println("No se aceptan decimales")
		}

		fmt.Println("Ingrese el valor del radio del circulo",i,": "); //1)
		_, _, _ = bufio.NewReader(os.Stdin).ReadLine()
		_, err = fmt.Scanln(&radio); // si err == nil entonces está bien
	}

	area, perimetro := calculaAreaPerimetro(radio);

	var funcion uint8


	fmt.Println("Que valor desea imprimir?")
	fmt.Println("1) Area")
	fmt.Println("2) Perimetro")
	fmt.Scanln(&funcion); 

	for funcion!=1 && funcion!=2 { // 4)
		fmt.Println(funcion,": Que valor desea imprimir?")
		fmt.Println("1) Area")
		fmt.Println("2) Perimetro")
		fmt.Scanln(&funcion); 
		fmt.Println("")
	}

	switch { // 7)
	case funcion==1:
		fmt.Println("El area es: ",area)
	case funcion==2:
		fmt.Println("El perimetro es: ",perimetro)
	default:
		fmt.Println("escoge una opcion correcta por favor")
	}

	fmt.Println( calculaAreaPerimetro(radio) ); // 1)

}