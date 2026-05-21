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

var tasks = []Task{}
var idCounter int

// Funciones

// Funcion para crear una tarea
func CreateTask() {
	var title, description string
	fmt.Print("Ingrese el titulo de la tarea: ")
	fmt.Scanln(&title)
	fmt.Print("Ingrese la descripcion de la tarea: ")
	fmt.Scanln(&description)

	idCounter++
	newTask := Task{
		ID: idCounter,
		Title: title,
		Description: description,
		Completed: false,
	}
	tasks = append(tasks, newTask)
	fmt.Println("Tarea creada exitosamente")
}

// Funcion para listar las tareas
func ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("No hay tareas creadas")
		return
	}
	fmt.Println("Lista de Tareas:")
	for _, task := range tasks {
		status := "Pendiente"
		if task.Completed {
			status = "Completada"
		}
		fmt.Printf("ID: %d, Titulo: %s, Descripcion: %s, Completada: %s\n", task.ID, task.Title, task.Description, status)
	}
}

// Funcion para actualizar una tarea
func UpdateTask() {
	var id int
	fmt.Print("Ingrese el ID de la tarea: ")
	fmt.Scanln(&id)
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = !task.Completed
			fmt.Println("Tarea actualizada exitosamente")
			return
		}
	}
	fmt.Println("Tarea no encontrada")
}
	
// Funcion para eliminar una tarea
func DeleteTask() {
	var id int
	fmt.Print("Ingrese el ID de la tarea: ")
	fmt.Scanln(&id)
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Tarea eliminada exitosamente")
			return
		}
	}
	fmt.Println("Tarea no encontrada")
}



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