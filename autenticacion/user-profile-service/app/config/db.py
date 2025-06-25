from pymongo import MongoClient
from dotenv import load_dotenv
import os

load_dotenv()

class MongoDB:
    _instancia = None

    def __new__(cls):
        if cls._instancia is None:
            cls._instancia = super(MongoDB, cls).__new__(cls)
            uri = os.getenv("MONGODB_URI")
            db_name = os.getenv("DB_NAME")
            cls._instancia.cliente = MongoClient(uri)
            cls._instancia.db = cls._instancia.cliente[db_name]
        return cls._instancia

    def obtener_db(self):
        return self.db
