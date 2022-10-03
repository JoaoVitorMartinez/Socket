# Utilização do Socket para comunicação TCP -  Golang

Esse projeto é destinado ao desenvolvimento de um socket em Golang, a comunicação entre cliente e servidor será realizada através de uma conexão TCP.

Através do socket o cliente pode enviar dados ao servidor (Seno, Coseno, Tangente) para que o mesmo efetue o cálculo da área de um triângulo qualquer.

## 🚀 Começando

Você pode consultar as informações abaixo para utilizar o projeto.

Consulte **Instalação** para saber como implantar o projeto.

## 📦 Pré-requisitos - Necessários para execução

  - Versão da linguagem Go 1.19.x
  

## 🔧 Instalação

  - Primeiramente é necessário efetuar a instalação da linguagem pelo **[Site](https://go.dev/dl/)**
  - Utilize uma IDE de sua preferência para executar os arquivos **servidor.go** e **cliente.go**;
  - Para execuação dos arquivos utilize a seguinte sintaxe no terminal: **go run <nomedoarquivo.go>**
  

## ⚙️ Executando os testes

  - Execute o arquivo **servidor.go**, você receberá a mensagem **"Servidor aguardando conexões..."**;
  - Execute o arquivo **cliente.go**, após isso você deve receber a mensagem **"Conexão aceita..."** isso indica que a conexão está funcionando;
  - Executando os arquivos a mensagem **"Entre com os valores do triângulo (Cateto oposto, adjacente e hipotetusa):"** será exibida na tela do cliente;
  - Feito isso você deve inserir os valores separados por espaço;
  - Após isso os valores serão enviar ao servidor e o resultado será retornado ao cliente informando a área do triângulo;


## 📋 Detalhando o código

### - Realizando a conexão entre servidor e cliente: 

**servidor.go**

func main() {

	fmt.Println("Servidor aguardando conexões...")

	var wg sync.WaitGroup
	chSeno := make(chan string)

	// ouvindo na porta 8081 via protocolo tcp/ip
	ln, erro1 := net.Listen("tcp", ":8081")
	if erro1 != nil {
		fmt.Println(erro1)
	}

	// aceitando conexões
	conexao, erro2 := ln.Accept()
	if erro2 != nil {
		fmt.Println(erro2)
		os.Exit(3)
	}

	defer ln.Close()

	fmt.Println("Conexão aceita...")


**cliente.go**

    conexao, erro1 := net.Dial("tcp", "127.0.0.1:8081")
       if erro1 != nil {
          fmt.Println(erro1)
          os.Exit(3)
      }
      
      
### - Aqui podemos verificar como o cliente.go captura as entradas sem o \r\n:

     for {
	 leitor := bufio.NewReader(os.Stdin)
	 fmt.Print("Entre com os valores do triângulo (Cateto oposto, adjacente e hipotetusa): ")
	 co, err := leitor.ReadString(' ')
	 ca, err := leitor.ReadString(' ')
	 h, err := leitor.ReadString(' ')
	 if err != nil {
		 fmt.Println(err)
		 os.Exit(3)
	 }
	
	
### - Aqui podemos ver como o servidor.go processa a mensagem recebida e efetua o cálculo utilizando as Go Rotines:

    for {
	mensagem, erro3 := bufio.NewReader(conexao).ReadString('\n')
	if erro3 != nil {
		fmt.Println(erro3)
		os.Exit(3)
        }
         

      		fmt.Print("Mensagem recebida:", valoresDoTriangulo)

		wg.Add(3)
		go func(wg *sync.WaitGroup) string {
			defer wg.Done()

			// chSeno <- seno(valoresDoTriangulo)
			time.Sleep(1 * time.Second)
			fmt.Println("fim go routine 1...")
			return "chSeno"
		}(&wg)

		go func(wg *sync.WaitGroup) string {
			defer wg.Done()
			// chTangente <- tangente(valoresDoTriangulo)
			fmt.Println("fim go routine 3...")

			return "chTangente"

		}(&wg)

		go func(wg *sync.WaitGroup) string {
			defer wg.Done()
			time.Sleep(3 * time.Second)
			
			
### - Nessa parte o servidor.go devolve o resultado ao cliente.go e o mesmo printa o resultado:

  **servidor.go**

     conexao.Write([]byte(string(mensagem + "\r\n")))
  
  **cliente.go**
  
     fmt.Println("Resposta do servidor: " + resp)
  
--- 

## 🛠️ Construído com

* [Golang](https://go.dev/)

## ✒️ Autores

* [**João Vitor**](https://github.com/JoaoVitorMartinez) - Desenvolvimento & Documentoção
* [**Rodrigo Gonçalves**](https://github.com/RodrigoSVG) - Desenvolvimento & Documentoção


---
