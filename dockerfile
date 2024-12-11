# Usa una imagen base de Go
FROM golang:1.19

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del proyecto al contenedor
COPY . .

# Descarga las dependencias de Go
RUN go mod tidy

# Construye la aplicación
RUN go build -o app .

# Expone el puerto en el que corre tu aplicación (por ejemplo, 8080)
EXPOSE 4000

# Comando para ejecutar la aplicación
CMD ["./app"]