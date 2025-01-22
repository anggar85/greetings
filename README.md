# Saludos en GO

Este paquete proporciona ima forma siomple de obtner saludos.

## Instalacion
Ejecuta:
```bash
go get -u github.com/anggar85/greetings
```

## Uso
Aqui tienes un ejemplo de como se usa

```go
package main

import (
    "fmt"
    "github.com/anggar85/greetings"
)

func main(){
	message, err := greetings.Hello("Emma")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
```