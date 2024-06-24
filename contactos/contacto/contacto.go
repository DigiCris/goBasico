package contacto

import "fmt"

//Contacto  5) acá hice la estructura publica
type Contacto struct {
	Name  string
	Phone string
	Email string
	edad uint8 // 4) debería ser una variable privada, no sé porque la logro leer
}

// 6) metodos de clases usando punteros
func (this *Contacto) SetName(value string) {
	this.Name = value
}

func (this *Contacto) SetPhone(value string) {
	this.Phone = value
}

func (this *Contacto) SetEmail(value string) {
	this.Email = value
}

func (this *Contacto) SetEdad(value uint8) {
	this.edad = value
}

func (this *Contacto) SetAll(Name,Phone,Email string, edad uint8) {
	this.Name = Name
	this.Phone = Phone
	this.Email = Email
	this.edad = edad
}

// 7) stringer
func (this * Contacto) String() string {
	return fmt.Sprintf("name: %s\nEmail: %s\nedad: %d\n", this.Name, this.Phone, this.Email, this.edad);
}