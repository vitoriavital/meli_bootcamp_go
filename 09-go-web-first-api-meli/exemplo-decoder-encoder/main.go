package main

import (
	"encoding/json"
	"log"
	"os"
)

type Aluno struct {
	Nome  string  `json:"name"`
	Idade int     `json:"idade"`
	Nota  float64 `json:"nota"`
}

func main() {
	alunos := []Aluno{
		{"João", 20, 8.5},
		{"Maria", 22, 9.0},
		{"João", 19, 7.5},
	}

	file, err := os.Create("alunos.json")
	if err != nil {
		log.Fatalf("Erro ao criar arquivo %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	//Escrever cada aluno como uma linha JSON
	for _, aluno := range alunos {
		err := encoder.Encode(aluno)
		if err != nil {
			log.Fatalf("Erro ao codificar aluno %v", err)
		}
	}
}
