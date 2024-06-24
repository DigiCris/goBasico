package main

import "fmt"

// Definimos la interfaz "Animal"
type Animal interface {
    HacerSonido()
}

// Definimos el tipo "Perro"
type Perro struct {}

// Definimos el método "HacerSonido()" para el tipo "Perro"
func (p Perro) HacerSonido() {
    fmt.Println("¡Guau!")
}

// Definimos el tipo "Gato"
type Gato struct {}

// Definimos el método "HacerSonido()" para el tipo "Gato"
func (g Gato) HacerSonido() {
    fmt.Println("¡Miau!")
}

// Esta función puede aceptar cualquier tipo que implemente la interfaz "Animal"
func ImprimirSonido(a Animal) {
    a.HacerSonido()
}

func main() {
    // Interfaces básicas

    // Creamos un objeto "Perro" y lo pasamos a la función ImprimirSonido()
    var perro Perro
    ImprimirSonido(perro) // Imprime "¡Guau!"

    // Creamos un objeto "Gato" y lo pasamos a la función ImprimirSonido()
    var gato Gato
    ImprimirSonido(gato) // Imprime "¡Miau!"



    // Lista de interfaces: tengo una lista de objetos que implementan cierta interfaz pero no necesito saber que son cada uno, es una forma de poder poner lista de cosas distintas tambien
    var animales []Animal
    animales = append(animales, Perro{}, Gato{}, Perro{}, Gato{})

    // Recorremos la lista y llamamos al método "HacerSonido()" de cada animal
    for _, animal := range animales {
        animal.HacerSonido()
    }
}
