# 🔐 Dominio: Autenticación - Ecomvida

Este dominio gestiona todo lo relacionado con el inicio de sesión, perfiles de usuario, direcciones, preferencias personalizadas y asignación de roles y permisos.

---

## 🧩 Microservicios Incluidos

| Microservicio             | Lenguaje   | BD             | Patrón         | Puerto |
|---------------------------|------------|----------------|----------------|--------|
| `auth-service`            | Node.js    | PostgreSQL     | Factory Method | 3001   |
| `user-profile-service`    | Python     | MongoDB Atlas  | Singleton      | 3002   |
| `address-service`         | Go         | MySQL          | Repository     | 3003   |
| `user-preferences-service`| Node.js    | DynamoDB Local | Observer       | 3004   |
| `role-permission-service` | Go         | PostgreSQL     | Strategy       | 3005   |

---

## 🧱 Características Generales

- Todos los microservicios tienen:
  - Endpoints REST separados por archivo.
  - Dockerización optimizada para producción.
  - Conexión a base de datos protegida por `.env`.
  - Patrón de diseño específico.
  - Estructura limpia, modular y desacoplada.

- No se requiere entorno local: cada microservicio está diseñado directamente para su despliegue en AWS u otra infraestructura cloud.

---

## 🚀 Flujo de Trabajo

1. Cada microservicio se valida de forma individual (lógica, pruebas, estructura).
2. Las pruebas se realizan vía Postman.
3. Todo el código está preparado para CI/CD con GitHub Actions.
4. Despliegue a EC2 con API Gateway, Load Balancer y Auto Scaling (una vez completado todo el sistema).

---

## 📦 Estructura del Dominio

autenticacion/
├── auth-service/
├── user-profile-service/
├── address-service/
├── user-preferences-service/
└── role-permission-service/

---

## ✅ Estado

✔️ Dominio autenticación finalizado y validado.  
Esperando integración con frontend y despliegue global al completar los demás dominios.