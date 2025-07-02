# Imagen base mínima
FROM node:18-alpine

# Directorio de trabajo
WORKDIR /app

# Copiar solo lo necesario y aprovechar cache
COPY package*.json ./
RUN npm install --omit=dev

# Copiar el resto del código fuente
COPY . .

# Asegurar permisos si se usan
RUN chmod -R 755 /app

# Variables de entorno (opcional)
ENV NODE_ENV=production

# Exponer el puerto
EXPOSE 3001

# Comando principal
CMD ["npm", "start"]
