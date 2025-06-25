from fastapi import HTTPException
from app.models.user_model import UserModel

def obtener_usuario(id: str):
    modelo = UserModel()
    usuario = modelo.obtener_por_id(id)
    if not usuario:
        raise HTTPException(status_code=404, detail="Usuario no encontrado")
    return usuario
