from fastapi import HTTPException
from app.models.user_model import UserModel
from app.schemas.user_schema import UserCreate

def crear_usuario(data: UserCreate):
    modelo = UserModel()
    user_dict = data.dict()
    user_id = modelo.crear(user_dict)
    return { "id": user_id, **user_dict }
