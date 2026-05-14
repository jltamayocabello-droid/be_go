// Colecciones de Datos

// Arrays

package main
import "fmt"

func main() { 
	
	var nombres [4]string
	nombres[0] = "Pedro"
	nombres[1] = "Juan"
	nombres[2] = "Maria"
	nombres[3] = "Luis"


	apellidos := [4]string{"Perez", "Gomez", "Garcia", "Lopez"}
	
	fmt.Println(nombres)
	fmt.Println(apellidos[0])
}