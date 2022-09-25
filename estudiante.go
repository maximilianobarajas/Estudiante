
package main
  
import (
    "fmt"
)
  
//Declaramos la estructura
type Estudiante struct{
      
    // defining struct variables
    nombre string
    apellido string
    edad int
    carrera string
}
  
//Funcion de descripción
func (estudiante Estudiante) imprimir_detalles(){
  
    fmt.Printf("Nombre completo: %s %s", estudiante.nombre, estudiante.apellido)
    fmt.Printf("\nCarrera: %s",estudiante.carrera)
    fmt.Printf("\nEdad: %d \n", estudiante.edad)
}

func main() {
    //Instanciamos la estructura
    estudiante1 := Estudiante{"Alan", "Perez", 22,"Lic. en Derecho"}
      
    //Imprimimos la estructura
    estudiante1.imprimir_detalles()
}


