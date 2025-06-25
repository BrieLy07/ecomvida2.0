from bson import ObjectId
from app.config.db import MongoDB

class UserModel:
    def __init__(self):
        self.db = MongoDB().obtener_db()
        self.coleccion = self.db["usuarios"]

    def crear(self, data):
        resultado = self.coleccion.insert_one(data)
        return str(resultado.inserted_id)

    def obtener_por_id(self, id):
        usuario = self.coleccion.find_one({"_id": ObjectId(id)})
        if usuario:
            usuario["_id"] = str(usuario["_id"])
        return usuario

    def actualizar(self, id, data):
        self.coleccion.update_one({"_id": ObjectId(id)}, {"$set": data})
        return self.obtener_por_id(id)

    def eliminar(self, id):
        resultado = self.coleccion.delete_one({"_id": ObjectId(id)})
        return resultado.deleted_count > 0
