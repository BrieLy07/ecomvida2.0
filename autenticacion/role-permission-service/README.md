# 🛡️ Role & Permission Service - Ecomvida

Microservicio para la gestión de roles, permisos y asignación de roles a usuarios en el sistema Ecomvida.

---

## 🧩 Descripción del Servicio

Este servicio permite crear y consultar roles y permisos, así como asignar roles a usuarios. Utiliza Go, PostgreSQL (RDS) y aplica el patrón de diseño **Strategy** para desacoplar la lógica de inserción.

---

## 🏗️ Estructura del Proyecto

role-permission-service/
├── config/ # Conexión Singleton a PostgreSQL
├── controllers/ # Un archivo por endpoint
├── models/ # Modelos de Rol, Permiso y UsuarioRol
├── routes/ # Definición de rutas
├── strategies/ # Estrategias por entidad (Strategy Pattern)
├── main.go # Inicialización de la app
├── .env.example
├── Dockerfile
├── go.mod
└── README.md

---

## 🧪 Endpoints Disponibles

| Método | Endpoint                 | Descripción                            |
|--------|--------------------------|----------------------------------------|
| GET    | `/roles`                | Listar roles disponibles               |
| POST   | `/roles`                | Crear nuevo rol                        |
| GET    | `/permissions`          | Listar permisos disponibles            |
| POST   | `/permissions`          | Crear nuevo permiso                    |
| GET    | `/users/:id/roles`      | Obtener roles asignados a un usuario   |
| POST   | `/users/:id/roles`      | Asignar roles a un usuario             |

---

## 🧱 Patrón de Diseño Utilizado

Se implementa el patrón **Strategy**, permitiendo ejecutar diferentes lógicas de inserción (rol, permiso) a través de una interfaz común (`Estrategia`).

---

## ⚙️ Variables de Entorno

Ver archivo `.env.example`:

PORT=3005
DB_HOST=role-db.cg9go1thnm7i.us-east-1.rds.amazonaws.com
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_clave
DB_NAME=role_service

---

## 🐳 Dockerización

El microservicio está empaquetado para producción. Compila la app en Go y expone el puerto `3005`.

---

## ✅ Pruebas

Pruebas manuales realizadas con **Postman**, verificando todos los endpoints y la lógica de asignación/consulta.