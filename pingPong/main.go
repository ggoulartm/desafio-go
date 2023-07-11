package main

import (
	"fmt"
	"time"
)

func main() {

	for {
		c1 := make(chan string)
		go func() { //Recebemos os valores "um"e depois "dois"conforme o esperado.
			time.Sleep(1 * time.Second) //Observe que o tempo total de execução é de apenas
			//~ 2 segundos,
			//pois o 1 e o 2 segundos são Sleeps executados simultaneamente.
			c1 <- "ping"
		}()

		c2 := make(chan string)
		go func() {

			time.Sleep(1 * time.Second)
			c2 <- "pong"
		}()
		for i := 0; i < 2; i++ {
			select { //Usaremos select para aguardar esses
			//dois valores simultaneamente, imprimindo cada um à medida que chegam.
			case msg1 := <-c1:
				fmt.Println(msg1)
				time.Sleep(1 * time.Second)
			case msg2 := <-c2:
				fmt.Println(msg2)
				time.Sleep(1200 * time.Microsecond)
			}
		}
	}
}
