from fastapi import HTTPException
from app.models.user_model import UserModel

def eliminar_usuario(id: str):
    modelo = UserModel()
    eliminado = modelo.eliminar(id)
    if not eliminado:
        raise HTTPException(status_code=404, detail="Usuario no encontrado")
    return { "mensaje": "Usuario eliminado correctamente" }
