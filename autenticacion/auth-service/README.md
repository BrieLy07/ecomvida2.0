# 🔐 Auth Service - Ecomvida

Microservicio de autenticación para el sistema Ecomvida.

---

## 🧩 Descripción del Servicio

Este servicio gestiona el registro, login, cierre de sesión, recuperación y renovación de credenciales de los usuarios. Utiliza JWT para la autenticación y sigue el patrón de diseño **Factory Method**.

---

## 🏗️ Estructura del Proyecto

auth-service/
├── src/
│ ├── controllers/ # Lógica de cada endpoint
│ ├── routes/ # Rutas expuestas
│ ├── services/ # Lógica de negocio con Factory
│ ├── models/ # Modelo User (Sequelize)
│ ├── config/ # Configuración de DB (PostgreSQL RDS)
│ ├── utils/ # Utilidades (ej: tokens)
│ └── app.js # Aplicación principal
├── Dockerfile
├── .env.example
├── .gitignore
├── package.json
└── README.md

---

## 🧪 Endpoints Disponibles

| Método | Endpoint             | Descripción                      |
|--------|----------------------|----------------------------------|
| POST   | `/register`          | Registro de nuevo usuario        |
| POST   | `/login`             | Inicio de sesión                 |
| POST   | `/logout`            | Cierre de sesión (blacklist JWT) |
| POST   | `/forgot-password`   | Solicitud de recuperación        |
| POST   | `/reset-password`    | Cambio de contraseña             |
| POST   | `/refresh-token`     | Obtener nuevo token JWT          |

---

## 🧱 Patrón de Diseño Utilizado

Se implementa el patrón **Factory Method** en el módulo `services/factory.js` para gestionar la lógica de los distintos procesos de autenticación de forma desacoplada.

---

## ⚙️ Variables de Entorno

Ver archivo `.env.example`:

---

## 🐳 Dockerización

El microservicio está preparado para producción. El `Dockerfile` instala solo dependencias necesarias (`--omit=dev`) y expone el puerto `3001`.

---

## ✅ Pruebas

Las pruebas de endpoints se realizan con **Postman**, y se incluye validación de respuesta esperada para cada caso.

---

## 🛡️ Seguridad

- Tokens JWT de acceso (1h) y refresh (7 días).
- Passwords cifradas con `bcrypt`.
- Blacklist de tokens para cerrar sesión.