package main

import (
	"contactos/contacto"
	"fmt"
)
//import "bufio"
//import "os"


/* Ejercicio

Crear un programa que maneje una agenda de contactos. La agenda debe contener la siguiente información de cada contacto:

	Nombre
	Teléfono
	Correo electrónico

Requisitos:

1) Array: Crear un array de 5 contactos precargados.
2) Slice: Crear un slice que permita agregar nuevos contactos a la agenda.
3) Map: Crear un mapa que permita buscar un contacto por su número de teléfono.
4) Estructura privada: Crear una estructura privada contacto que contenga los datos de cada contacto.
5) Estructura pública: Crear una estructura pública Agenda que contenga el array y el slice de contactos.
6) Función para modificar elemento de estructura por puntero: Crear una función que permita modificar el teléfono de un contacto en la agenda, utilizando punteros.
7) Stringer: Implementar la interfaz fmt.Stringer en la estructura contacto para que se pueda imprimir la información de un contacto de forma amigable.

Pasos a seguir:

1) Crear la estructura contacto con los campos nombre, teléfono y correo electrónico.
2) Crear la estructura Agenda que contenga el array y el slice de contactos.
3) Implementar la interfaz fmt.Stringer en la estructura contacto.
4) Crear funciones para:
	a)Agregar un nuevo contacto al slice.
	b)Buscar un contacto por teléfono en el mapa.
	c)Modificar el teléfono de un contacto en la agenda.
5) Crear el programa principal que demuestre el uso de todos los elementos solicitados.
*/


func main(){
	defer fmt.Println("gracias por usar nuestra agenda");
	contactByPhone := make(map[string]string);

	// 1)
	var favoritos [5]contacto.Contacto;
	for i:=0; i<5; i++ {
		favoritos[i].SetAll("cris","1166","a@b.c",34)
		contactByPhone[favoritos[i].Name] = favoritos[i].Phone;
	}
	fmt.Println(favoritos[0].String())

	// 2)
	agenda := []contacto.Contacto{}; // 5
	agenda = append(agenda, favoritos[:]...) // solo puedo hacer un append de slices y por eso tengo que hacerle un slice a favoritos de todos sus valores
	fmt.Println(agenda)
	name := ""
	phone := ""
	for name!="-1" { // me mantengo en loop agregando gente a la agenda
		name, phone = ingresarContacto(&agenda)
		if name!="-1" {
			contactByPhone[name] = phone;
		}
	}

	// 3)

	fmt.Println("Que contacto desea buscar?");
	fmt.Scanln(&name);
	fmt.Printf("el telefono de %s es %s ",name,contactByPhone[name]);
	fmt.Println("");


	// Extra) Acá estoy imprimiendo todo el mapping usando range
	fmt.Println("");
	fmt.Println("Ahora imprimo todo el mapping");
	for i, v := range contactByPhone {
		fmt.Println(i,v);
	}

	// 4) puse edad privada pero no sé porque la logro leer


	
}

// 6) funciones con punteros de referencias
func ingresarContacto(agenda *[]contacto.Contacto) (Name, Phone string){
	add := "n"
	var newContact contacto.Contacto;
	fmt.Print("Desea ingresar un nuevo contacto? y/n: ");
	fmt.Scanln(&add)
	fmt.Println("");
	for add!="y" && add!="n" {
		fmt.Print("Desea ingresar un nuevo contacto? y/n: ");
		fmt.Scanln(&add)
		fmt.Println("");
	}
	if add=="y" {
		fmt.Print("Ingrese nombre: ");
		fmt.Scanln(&newContact.Name)
		fmt.Println("");
		fmt.Print("Ingrese telefono: ");
		fmt.Scanln(&newContact.Phone)
		fmt.Println("");
		fmt.Print("Ingrese Email: ");
		fmt.Scanln(&newContact.Email)
		fmt.Println("");
		fmt.Print("Ingrese Edad: ");
		var edad uint8;
		fmt.Scanln(&edad)
		newContact.SetEdad(edad)
		fmt.Println("");
		*agenda = append(*agenda,newContact)
		fmt.Println(*agenda);		
		return newContact.Name,newContact.Phone
	} else{
		return "-1","-1";
	}

}