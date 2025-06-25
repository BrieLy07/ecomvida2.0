# Imagen base
FROM python:3.11-slim

# Establecer directorio de trabajo
WORKDIR /app

# Copiar e instalar dependencias
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copiar el resto del código
COPY ./app ./app
COPY .env .env

# Puerto expuesto
EXPOSE 3002

# Comando para ejecutar el servidor
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "3002"]
