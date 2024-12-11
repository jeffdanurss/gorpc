package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Define un tipo que implementará el servicio RPC
type HelloService struct{}

// Define un método que se expondrá a través de RPC
func (h *HelloService) SayHello(request string, reply *string) error {
	*reply = "Hola, " + request
	return nil
}

func main() {
	// Crea una instancia de HelloService
	helloService := new(HelloService)

	// Registra el servicio RPC
	err := rpc.Register(helloService)
	if err != nil {
		fmt.Println("Error registrando el servicio:", err)
		return
	}

	// Escucha en el puerto 1234
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Servidor RPC escuchando en :1234...")

	// Acepta las conexiones entrantes y sirve las peticiones RPC
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error al aceptar la conexión:", err)
			continue
		}

		// Manejamos la conexión con el servicio RPC
		go rpc.ServeConn(conn)
	}
}
