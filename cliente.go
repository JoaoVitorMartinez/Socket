package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

var (
	conn net.Conn
	err  error
)

func main() {

	// conectando na porta 8081 via protocolo tcp/ip na máquina local
	// fica tentando se conectar em caso de falha e que o servidor não foi parado de forma abrupta
	for {
		fmt.Println("Conectando ao servidor...")
		conn, err = net.Dial("tcp", "127.0.0.1:8081")
		if err == nil {
			break
		}
		fmt.Println(err)
		time.Sleep(time.Second * 1)
	}
	fmt.Println("Conexão aceita.")

	for {
		// lendo entrada do terminal
		leitor := bufio.NewReader(os.Stdin)
		fmt.Print("Entre com os valores do triângulo (Cateto oposto, adjacente e hipotetusa): ")
		co, err := leitor.ReadString(' ')
		ca, err := leitor.ReadString(' ')
		h, err := leitor.ReadString(' ')
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}

		// escrevendo a mensagem na conexão (socket)
		fmt.Fprintf(conn, co+ca+h+"\n")

		// ouvindo a resposta do servidor (eco)
		resp, err3 := bufio.NewReader(conn).ReadString(' ')
		if err3 != nil {
			fmt.Println(err3)
			os.Exit(3)
		}
		// escrevendo a resposta do servidor no terminal
		fmt.Println("Resposta do servidor: " + resp)
	}
}
