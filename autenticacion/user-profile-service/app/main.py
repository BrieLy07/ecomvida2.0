from fastapi import FastAPI
from dotenv import load_dotenv
import os

# Importar controladores
from app.controllers.create_user import crear_usuario
from app.controllers.get_user import obtener_usuario
from app.controllers.update_user import actualizar_usuario
from app.controllers.delete_user import eliminar_usuario
from app.schemas.user_schema import UserCreate, UserUpdate

load_dotenv()

app = FastAPI()

# Rutas
@app.post("/users")
def crear(data: UserCreate):
    return crear_usuario(data)

@app.get("/users/{id}")
def obtener(id: str):
    return obtener_usuario(id)

@app.put("/users/{id}")
def actualizar(id: str, data: UserUpdate):
    return actualizar_usuario(id, data)

@app.delete("/users/{id}")
def eliminar(id: str):
    return eliminar_usuario(id)
