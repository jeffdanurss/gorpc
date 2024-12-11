package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// Conectar al servidor RPC en la dirección localhost:1234
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error al conectar al servidor:", err)
		return
	}
	defer client.Close()

	// Definir el nombre que se enviará al servidor
	name := "Juan"
	var reply string

	// Realizar la llamada RPC a la función SayHello
	err = client.Call("HelloService.SayHello", name, &reply)
	if err != nil {
		fmt.Println("Error al hacer la llamada RPC:", err)
		return
	}

	// Imprimir el saludo recibido
	fmt.Println("Respuesta del servidor:", reply)
}
