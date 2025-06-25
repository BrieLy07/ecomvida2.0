# 👤 User Profile Service - Ecomvida

Microservicio para la gestión de perfiles de usuario dentro del sistema Ecomvida.

---

## 🧩 Descripción del Servicio

Este servicio permite crear, consultar, actualizar y eliminar perfiles de usuario. Utiliza FastAPI y MongoDB Atlas como base de datos. Aplica el patrón de diseño **Singleton** para manejar una única instancia de conexión con la base de datos.

---

## 🏗️ Estructura del Proyecto

user-profile-service/
├── app/
│ ├── config/ # Conexión MongoDB (Singleton)
│ ├── controllers/ # Un archivo por endpoint
│ ├── models/ # Lógica de acceso a MongoDB
│ ├── schemas/ # Validaciones con Pydantic
│ └── main.py # Servidor FastAPI
├── .env.example
├── .gitignore
├── Dockerfile
├── requirements.txt
└── README.md

---

## 🧪 Endpoints Disponibles

| Método | Endpoint       | Descripción                     |
|--------|----------------|----------------------------------|
| POST   | `/users`       | Crear perfil de usuario         |
| GET    | `/users/:id`   | Obtener perfil por ID           |
| PUT    | `/users/:id`   | Actualizar perfil existente     |
| DELETE | `/users/:id`   | Eliminar perfil de usuario      |

---

## 🧱 Patrón de Diseño Utilizado

Se utiliza el patrón **Singleton** para la conexión a MongoDB, garantizando una única instancia compartida por toda la aplicación.

---

## ⚙️ Variables de Entorno

Ver archivo `.env.example`:

---

## 🐳 Dockerización

El microservicio está empaquetado para producción usando un contenedor ligero de Python. Expone el puerto `3002` y ejecuta `uvicorn` como servidor principal.

---

## ✅ Pruebas

Las pruebas se realizan con **Postman**, validando los endpoints individuales de forma separada.