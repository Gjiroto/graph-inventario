# Usar una imagen base con Go preinstalado
FROM golang:1.20-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el archivo go.mod y go.sum y descargar las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente del proyecto
COPY . .

# Copiar el código fuente del proyecto
COPY . .

# Definir el comando por defecto para ejecutar la aplicación
CMD ["go", "run", "server.go"]
