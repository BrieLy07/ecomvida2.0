# Imagen base de Go
FROM golang:1.20-alpine

# Crear carpeta de trabajo
WORKDIR /app

# Copiar archivos necesarios
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copiar el resto del código
COPY . .

# Compilar
RUN go build -o main .

# Exponer puerto
EXPOSE 3003

# Comando al iniciar contenedor
CMD ["./main"]
