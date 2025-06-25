from fastapi import HTTPException
from app.models.user_model import UserModel
from app.schemas.user_schema import UserUpdate

def actualizar_usuario(id: str, data: UserUpdate):
    modelo = UserModel()
    actualizado = modelo.actualizar(id, data.dict(exclude_unset=True))
    if not actualizado:
        raise HTTPException(status_code=404, detail="Usuario no encontrado")
    return actualizado
