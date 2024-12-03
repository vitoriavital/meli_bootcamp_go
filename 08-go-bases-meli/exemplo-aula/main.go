package main

import (
	"fmt"
	"os"
)

func main() {
	// 1: criar arquivo
	file, err := os.Create("dados.txt")
	if err != nil {
		fmt.Println("Error ao criar o arquivo: ", err)
		return
	}
	defer file.Close() //fechar arquivo no final

	// 2: escrever dados no arquivo
	_, err = file.WriteString("Olá, este é o meu primeiro arquivo em Go!\n")
	if err != nil {
		fmt.Println("Error ao escrever no arquivo: ", err)
		return
	}
	_, err = file.WriteString("Manipular arquivos é mais fácil com Go\n")
	if err != nil {
		fmt.Println("Error ao escrever no arquivo: ", err)
		return
	}
	fmt.Println("Dados inseridos com sucesso!")

	// 3: abrir e ler o arquivo
	data, err := os.ReadFile("dados.txt")
	if err != nil {
		fmt.Println("Error ao ler o arquivo: ", err)
		return
	}

	//4: exibir os dados no console
	fmt.Println("\nConteúdo do arquivo:")
	fmt.Println(string(data))

}