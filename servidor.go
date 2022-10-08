package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Triangulo struct {
	catetoOposto    string
	catetoAdjacente string
	hipotenusa      string
}

// Função Seno
func seno(triangulo Triangulo) string {
	co, _ := strconv.ParseFloat(triangulo.catetoOposto, 64)
	hi, _ := strconv.ParseFloat(triangulo.hipotenusa, 64)

	seno := co / hi
	senoConvertido := strconv.FormatFloat(seno, 'f', -1, 64)
	return "\nSeno:" + senoConvertido
}

// Função Coseno
func coseno(triangulo Triangulo) string {
	ca, _ := strconv.ParseFloat(triangulo.catetoAdjacente, 64)
	hi, _ := strconv.ParseFloat(triangulo.hipotenusa, 64)

	coseno := ca / hi
	cosenoConvertido := strconv.FormatFloat(coseno, 'f', -1, 64)
	return "\nCoseno:" + cosenoConvertido
}

// Função Tangente
func tangente(triangulo Triangulo) string {
	co, _ := strconv.ParseFloat(triangulo.catetoOposto, 64)
	ca, _ := strconv.ParseFloat(triangulo.catetoAdjacente, 64)

	tangente := co / ca
	tangenteConvertida := strconv.FormatFloat(tangente, 'f', -1, 64)
	return "\nTangente:" + tangenteConvertida

}

func main() {
	var (
		ResultadoSeno     string
		ResultadoCoseno   string
		ResultadoTangente string
		wg                sync.WaitGroup
	)

	fmt.Println("Servidor aguardando conexões...")

	// ouvindo na porta 8081 via protocolo tcp/ip
	ln, erro1 := net.Listen("tcp", ":8081")
	if erro1 != nil {
		fmt.Println(erro1)
		/* Neste nosso exemplo vamos convencionar que a saída 3 está reservada para erros de conexão.
		IMPORTANTE: defers não serão executados quando utilizamos os.Exit() e a saída será imediata */
		os.Exit(3)
	}

	// aceitando conexões
	conexao, erro2 := ln.Accept()
	if erro2 != nil {
		fmt.Println(erro2)
		os.Exit(3)
	}

	defer ln.Close()

	fmt.Println("Conexão aceita...")
	// rodando loop contínuo (até que ctrl-c seja acionado)
	for {
		// Assim que receber o controle de nova linha (\n), processa a mensagem recebida
		mensagem, erro3 := bufio.NewReader(conexao).ReadString('\n')
		if erro3 != nil {
			fmt.Println(erro3)
			os.Exit(3)
		}

		valoresDoTriangulo := strings.Split(mensagem, " ")

		fmt.Print("Mensagem recebida:", valoresDoTriangulo)

		triangulo := Triangulo{
			catetoOposto:    valoresDoTriangulo[0],
			catetoAdjacente: valoresDoTriangulo[1],
			hipotenusa:      valoresDoTriangulo[2],
		}

		wg.Add(3)
		go func(wg *sync.WaitGroup) string {
			defer wg.Done()

			ResultadoSeno = seno(triangulo)

			time.Sleep(1 * time.Second)

			fmt.Println("fim go routine 1...")
			return "chSeno"
		}(&wg)

		go func(wg *sync.WaitGroup) string {
			defer wg.Done()
			ResultadoTangente = tangente(triangulo)
			fmt.Println("fim go routine 3...")

			return "chTangente"

		}(&wg)

		go func(wg *sync.WaitGroup) string {
			defer wg.Done()
			time.Sleep(3 * time.Second)

			ResultadoCoseno = coseno(triangulo)

			fmt.Println("fim go routine 2...")
			return "chCoseno"

		}(&wg)
		fmt.Println("Aguardando...")

		wg.Wait()

		fmt.Println("Fim...")

		conexao.Write([]byte(string(ResultadoSeno + ResultadoCoseno + ResultadoTangente + " ")))

	}
}
