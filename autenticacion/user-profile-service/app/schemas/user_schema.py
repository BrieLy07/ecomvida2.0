from pydantic import BaseModel, EmailStr, Field
from typing import Optional

class UserCreate(BaseModel):
    nombre: str
    apellido: str
    usuario: str
    correo: EmailStr
    numero: str

class UserUpdate(BaseModel):
    nombre: Optional[str]
    apellido: Optional[str]
    usuario: Optional[str]
    correo: Optional[EmailStr]
    numero: Optional[str]

class UserResponse(BaseModel):
    id: str = Field(..., alias="_id")
    nombre: str
    apellido: str
    usuario: str
    correo: EmailStr
    numero: str
