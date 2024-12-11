# Usa la imagen oficial de Go como base
FROM golang:1.19-alpine

# Define el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos necesarios al contenedor
COPY . .

# Instala las dependencias y compila el código Go
RUN go mod tidy && go build -o server .

# Expone el puerto 50051 (o el que utilices)
EXPOSE 50051

# Comando para ejecutar la aplicación
CMD ["./server"]
