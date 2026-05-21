package main

import (
	"fmt"
)

// Estructura de Tarea
type Task struct {
	ID int
	Title string
	Description string
	Completed bool
}

// Variables globales

var tasks = []Task
var idCounter int

// Funciones


// Funcion principal

func main() {
	for {
		fmt.Println("Task Manager")
		fmt.Println("1. Crear Tarea")
		fmt.Println("2. Listar Tareas")
		fmt.Println("3. Actualizar Tarea")
		fmt.Println("4. Eliminar Tarea")
		fmt.Println("5. Salir")
		fmt.Print("Seleccione una opción: ")


		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			CreateTask()
		case 2:
			ListTasks()
		case 3:
			UpdateTask()
		case 4:
			DeleteTask()
		case 5:
			fmt.Println("Gracias por usar el Task Manager")
			return
		default: 
			fmt.Println("Opcion no valida")
		}
	}
}