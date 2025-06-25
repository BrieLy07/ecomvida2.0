# Imagen base mínima de Node.js
FROM node:18-alpine

# Directorio de trabajo
WORKDIR /app

# Copiar solo archivos necesarios
COPY package*.json ./
RUN npm install --omit=dev

# Copiar el código fuente
COPY src/ ./src/
COPY .env .env

# Exponer el puerto de servicio
EXPOSE 3001

# Comando para producción
CMD ["npm", "start"]
