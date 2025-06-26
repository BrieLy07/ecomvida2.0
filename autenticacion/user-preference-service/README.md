# ⚙️ User Preferences Service - Ecomvida

Microservicio para la gestión de preferencias de usuario en Ecomvida.

---

## 🧩 Descripción del Servicio

Este servicio permite obtener y actualizar las preferencias personales de un usuario (tema, notificaciones, idioma). Utiliza Node.js, DynamoDB (modo local) y aplica el patrón de diseño **Observer** para simular notificaciones internas ante cambios.

---

## 🏗️ Estructura del Proyecto

user-preferences-service/
├── src/
│ ├── config/ # Conexión a DynamoDB local
│ ├── controllers/ # Un archivo por endpoint
│ ├── models/ # Modelo de Preferencia
│ ├── routes/ # Definición de rutas
│ ├── services/ # Lógica de negocio + Observer
│ └── app.js # Servidor Express
├── .env.example
├── Dockerfile
├── package.json
├── docker-compose.yml # DynamoDB local
└── README.md

---

## 🧪 Endpoints Disponibles

| Método | Endpoint                       | Descripción                       |
|--------|--------------------------------|-----------------------------------|
| GET    | `/users/:id/preferences`       | Obtener preferencias de usuario   |
| PUT    | `/users/:id/preferences`       | Actualizar preferencias del usuario|

---

## 🧱 Patrón de Diseño Utilizado

Se aplica el patrón **Observer** para simular suscriptores internos al momento de actualizar las preferencias. El sujeto notifica a los observadores registrados (en este caso, un log).

---

## ⚙️ Variables de Entorno

Ver archivo `.env.example`:

PORT=3004
AWS_REGION=us-east-1
DYNAMO_ENDPOINT=http://localhost:8000
DYNAMO_TABLE=preferencias

---

## 🐳 Dockerización

El microservicio está empaquetado para producción. DynamoDB se ejecuta de forma local usando `docker-compose`.

---

## ✅ Pruebas

Se realizan pruebas con **Postman**, conectándose a DynamoDB local. El código está listo para adaptarse a AWS cambiando solo las variables de entorno.