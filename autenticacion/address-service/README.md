# 📍 Address Service - Ecomvida

Microservicio para la gestión de direcciones de los usuarios dentro del sistema Ecomvida.

---

## 🧩 Descripción del Servicio

Este servicio permite crear, listar, actualizar y eliminar direcciones asociadas a un usuario. Utiliza Go, MySQL (RDS) y aplica el patrón de diseño **Repository** para separar la lógica de acceso a datos.

---

## 🏗️ Estructura del Proyecto

address-service/
├── config/ # Conexión a MySQL (Singleton)
├── controllers/ # Un archivo por endpoint
├── models/ # Estructura del modelo Address
├── repository/ # Interface y repositorio MySQL
├── routes/ # Definición de rutas
├── main.go # Aplicación principal
├── .env.example
├── Dockerfile
├── go.mod
├── README.md

---

## 🧪 Endpoints Disponibles

| Método | Endpoint                     | Descripción                      |
|--------|------------------------------|----------------------------------|
| POST   | `/users/:id/addresses`       | Crear nueva dirección            |
| GET    | `/users/:id/addresses`       | Listar direcciones del usuario   |
| PUT    | `/addresses/:addressId`      | Actualizar dirección             |
| DELETE | `/addresses/:addressId`      | Eliminar dirección               |

---

## 🧱 Patrón de Diseño Utilizado

Se implementa el patrón **Repository**, desacoplando la lógica del acceso a datos mediante una interfaz que puede ser sustituida fácilmente en el futuro.

---

## ⚙️ Variables de Entorno

Ver archivo `.env.example`:

---

## 🐳 Dockerización

El servicio está empaquetado para producción. El contenedor compila la aplicación en Go y expone el puerto `3003`.

---

## ✅ Pruebas

Las pruebas de los endpoints se realizan con **Postman**, validando las respuestas esperadas y el manejo de errores.