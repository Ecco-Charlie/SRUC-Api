# SRUC *Api*

> **Warning**
> Este proyecto actualmente se encuenta **Descontinuado**

## ¿Qué es SRUC?

SRUC (Sistema de Registro de Uso de Computadoras) es un proyecto que nace gracias a la falta de seguimiento de equipos de computo en la universidad.

Se tiene planeado que sea un sistema parecido a al de un Cibercafe (La computadora no puede ser accedida a menos que el usuario **(Alumno)** introduzca su matrícula), y darle seguimiento a que programas accede, y almacenar el que haya sido más usado, para posteriormente a fin de cada mes, generar un reporte con los datos de cada **Sala de Cómputo**.

## ¿Qué funcion tiene este repositorio?
Este repositorio almacena una API (Application Programming Interface) que nos ayuda a:
* Administrar la Base de Datos (MySQL).
* Tener un Dashboard para ver los datos de la BD y modificarlos.
* Ser una REST API para proporcionar los datos necesarios al **Cliente de escritorio (Consta de un demonio y una aplicación de escritorio)**.

## Tecnologías

### HTTP
Esta API esta echa en Golang, es un proyecto que me ayudo a mejorar mis habilidades y reforzar mis conocimientos creando API’s, previamente había echo API’s en **Spring Boot (Java)**, pero gracias a este proyecto, donde, no se ocuparon Frameworks Web (Gin,Gorilla Mux, etc), todo fue echo con las herramientas HTTP que nos proporciona el lenguaje.

### Autenticación
Fue implementado **JWT *(JSON Web Tokens)*** para tener una autentificación del **cliente *(Demonio)*** con la **API**.

### Base de Datos
Con la ayuda de un **ORM *(Object Relational Mapping)***, en este caso **GORM** se logro administrar la Base de Datos, simplificando hasta cierto punto su uso.

> Creado por *soft.exe*
