# Usa la imagen base de Go 1.22
FROM golang:1.22-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia solo los archivos go.mod y go.sum para aprovechar la caché de Docker
COPY go.mod go.sum ./

# Instala las dependencias de Go (en este caso, Fiber y otras que hayas definido)
RUN go mod tidy

# Copia todo el código fuente del proyecto dentro del contenedor
COPY . .

# Compila la aplicación Go con Fiber
RUN go build -o app .

# Expone el puerto que la aplicación va a usar (puedes cambiar el puerto si es necesario)
EXPOSE 8080

# Comando para ejecutar la aplicación Go con Fiber
CMD ["./app"]
