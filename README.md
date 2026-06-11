![Estado del Proyecto](https://img.shields.io/badge/Estado-Completado-green)
![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)
![Udemy Course](https://img.shields.io/badge/Curso-Backend%20Developer%20(Udemy)-ec1c24?logo=udemy&logoColor=white)

---

## 📖 Descripción del Proyecto
Este repositorio reúne el conjunto de prácticas, ejercicios y desarrollos realizados a lo largo de mi trayectoria de aprendizaje en el lenguaje de programación **Go (Golang)**. Todo el contenido forma parte de mi especialización académica orientada al desarrollo backend:

1. **Curso de Backend Developer (Udemy):** Un recorrido integral que abarca desde los fundamentos sintácticos del lenguaje, control de flujo y manipulación de datos, hasta conceptos avanzados como la gestión de memoria mediante punteros, composición orientada a objetos (interfaces/structs), persistencia local en archivos y concurrencia nativa (Goroutines/Channels).
2. **Proyecto Integrador (Task Manager):** Un gestor de tareas de consola interactivo ubicado en la carpeta [taskmanager](./taskmanager), desarrollado para aplicar de forma práctica los conceptos de slices, structs, funciones y flujos lógicos en un entorno real.

---

## 🎯 Objetivo
Consolidar el dominio técnico del lenguaje Go bajo estándares profesionales de desarrollo de software para el backend, logrando:
- Diseñar soluciones modulares de alto rendimiento utilizando los principios de **composición sobre herencia** característicos de Go.
- Integrar la **concurrencia nativa (modelo CSP)** de forma segura y eficiente para optimizar el rendimiento de procesos en segundo plano.
- Implementar arquitecturas de software robustas basadas en el control explícito de errores y la persistencia local de datos en formatos estándar como JSON.

---

## 🛠️ Requerimientos Técnicos / Temas Cubiertos
Este proyecto cumple con los estándares exigidos para el aprendizaje integral del desarrollo backend en Go:

### 1. Fundamentos & Lógica de Programación
- ✅ **Sintaxis y Tipos de Datos:** Declaración de variables (corta e implícita), constantes, tipos de datos primitivos y operadores en `a-introduccion`.
- ✅ **Estructuras de Control:** Toma de decisiones con `if/else`, condicionales múltiples con `switch` y bucles lógicos con `for` (único bucle de Go para bucles tradicionales, infinitos o iteradores de colecciones) en `b-intermedia`.
- ✅ **Modularidad y Funciones:** Declaración de funciones, paso de parámetros, múltiples valores de retorno, funciones anónimas, closures y el uso estratégico de `defer` para la liberación de recursos en `c-funciones`.

### 2. Estructuras de Datos Avanzadas & Gestión de Memoria
- ✅ **Punteros:** Manejo explícito de referencias de memoria, desreferenciación y paso por valor vs. paso por referencia en `d-punteros`.
- ✅ **Colecciones y Structs:** Creación de registros personalizados mediante `structs`, métodos asociados (`value and pointer receivers`) y manipulación avanzada de arreglos y slices en `e-estructuras` e `i-oop`.
- ✅ **Polimorfismo e Interfaces:** Definición de contratos mediante `interfaces`, uso de la interfaz vacía (`interface{}` / `any`) para el manejo de tipos dinámicos y desacoplamiento de componentes en `f-interfaces`.

### 3. Concurrencia, Errores y Persistencia
- ✅ **Concurrencia Nativa (CSP):** Ejecución asíncrona ligera mediante **Goroutines**, sincronización con **WaitGroups**, comunicación segura a través de **Channels** y multiplexación de canales utilizando la sentencia **Select** en `g-concurrencia`.
- ✅ **Manejo de Errores Idiomático:** Retornos explícitos de la interfaz predefinida `error`, creación de fallas personalizadas y la implementación segura de la recuperación ante pánicos (`panic` y `recover`) en `h-errores`.
- ✅ **Persistencia e Interoperabilidad:** Operaciones básicas de lectura y escritura de archivos locales en `j-archivos`, junto con la serialización y deserialización de estructuras complejas a formato **JSON** (`Marshal` y `Unmarshal`) en `k-json`.
- ✅ **Gestión de Paquetes (Módulos):** Estructuración de proyectos locales mediante **Go Modules**, importación de paquetes propios y control de visibilidad (variables y funciones exportadas/privadas) en `l-proyecto_paquetes` y `taskmanager`.

---

## 📂 Estructura del Proyecto

```
be_go/ (Go/)
│
├── ARCHIVOS/                             # Carpeta con todas las guías de estudio por temas
│   ├── a-introduccion/                   # Fundamentos sintácticos y tipos de datos (basico.go)
│   ├── b-intermedia/                     # Estructuras de control y flujo lógico
│   ├── c-funciones/                      # Modularidad, closures y uso de defer
│   ├── d-punteros/                       # Referencias de memoria y paso por dirección
│   ├── e-estructuras/                    # Agrupación de datos personalizados mediante structs
│   ├── f-interfaces/                     # Declaración de contratos y desacoplamiento de código
│   ├── g-concurrencia/                   # Goroutines, Channels, Select y WaitGroups
│   ├── h-errores/                        # Control defensivo y manejo de pánicos (panic/recover)
│   ├── i-oop/                            # Composición de structs y simulación de POO
│   ├── j-archivos/                       # Operaciones de lectura y escritura en disco
│   ├── k-json/                           # Serialización y deserialización de datos en formato JSON
│   └── l-proyecto_paquetes/              # Estructuración y visibilidad de paquetes modulares
│
└── taskmanager/                          # 🌟 PROYECTO PRÁCTICO: GESTOR DE TAREAS POR CONSOLA
    ├── go.mod                            # Archivo de definición del módulo de Go
    └── main.go                           # Código fuente del CRUD interactivo en terminal
```

---

## 🚀 Instrucciones de Ejecución

Para clonar y ejecutar cualquiera de los programas de este repositorio, asegúrate de tener instalado **Go (versión 1.22 o superior)** en tu sistema.

### 1. Clonar el repositorio
```bash
git clone https://github.com/jltamayocabello-droid/be_go.git
cd be_go
```

### 2. Ejecutar Ejercicios Temáticos
Puedes ejecutar directamente cualquier archivo de la carpeta `ARCHIVOS/` utilizando el comando `go run`:
```bash
# Ejemplo: Ejecutar el módulo de introducción básica
go run ARCHIVOS/a-introduccion/basico.go
```

### 3. Ejecutar el Gestor de Tareas (Task Manager)
El proyecto interactivo de consola está aislado en su propio módulo:
```bash
# Navegar al directorio del proyecto
cd taskmanager

# Ejecutar el gestor de tareas
go run main.go
```

---

## 🛠️ Decisiones Técnicas y Justificaciones

### Composición sobre Herencia (Foco en Struct Embedding)
**Decisión:** Evitar la complejidad de las jerarquías de clases tradicionales mediante el uso de la composición nativa de Go (acoplamiento de estructuras).
**Justificación:** Go no posee herencia basada en clases. La composición de structs e interfaces permite construir componentes altamente reutilizables y extensibles con menor acoplamiento, facilitando las pruebas unitarias y el mantenimiento.

### Concurrencia Basada en el Modelo CSP (Goroutines & Channels)
**Decisión:** Implementar la concurrencia nativa de Go para realizar tareas en paralelo, comunicando los procesos mediante canales en lugar de bloquear y compartir memoria.
**Justificación:** Reduce drásticamente las condiciones de carrera (*race conditions*) y la complejidad del desarrollo multihilo clásico, logrando un backend escalable con un consumo mínimo de recursos del sistema.

### Manejo Explícito de Errores
**Decisión:** Utilizar el retorno explícito del tipo `error` como el último argumento de las funciones, forzando la evaluación defensiva inmediata del error.
**Justificación:** Aumenta la robustez y legibilidad del software, previniendo fallos invisibles que en otros lenguajes se propagarían de manera silenciosa hasta romper el ciclo de vida de la aplicación.

---

## 📱 Áreas Técnicas Destacadas

| Área Técnica | Conceptos Clave | Descripción |
| :--- | :--- | :--- |
| 🐹 **Backend Core** | Sintaxis Go, Tipado Estricto | Compilación súper rápida y tipado estático robusto para servicios altamente confiables. |
| 🧬 **Concurrencia Nativa** | Goroutines, Channels, Select | Programación concurrente de alto rendimiento utilizando hilos ligeros administrados por el runtime de Go. |
| 🧱 **Arquitectura de Composición** | Structs, Interfaces, Embedding | Desarrollo desacoplado y orientado a interfaces que favorece la inyección de dependencias. |
| 📂 **Manipulación de Archivos** | `os`, `bufio`, `io` | Persistencia y almacenamiento directo en el disco local mediante la biblioteca estándar de Go. |
| 🌐 **Intercambio de Datos (JSON)** | `Marshal`, `Unmarshal` | Integración nativa rápida y eficiente para el parseo de datos en servicios web y APIs REST. |
| 📦 **Modularidad del Sistema** | Go Modules (`go.mod`) | Gestión interna y modularización bajo estándares modernos de desarrollo de la comunidad Go. |

---

## ✒️ Autor
**Jorge Tamayo Cabello**
*Diseñador Front-End*

---

## 📄 Licencia
Este repositorio es de carácter estrictamente académico y educativo. Todo el contenido es libre de ser consultado con fines de aprendizaje y referencia técnica.

---

## 🙏 Agradecimientos
- A **Udemy** por la sólida formación técnica brindada a través del curso de Backend Developer.
- A la **comunidad de Go** por diseñar y mantener un lenguaje tan ágil, claro y potente para el desarrollo backend.
- A todos los desarrolladores que promueven el ecosistema open-source facilitando recursos educativos de alta calidad.
