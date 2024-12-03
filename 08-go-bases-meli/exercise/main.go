package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/google/uuid"
)

type Student struct {
	Matricula	string
	Nome		string
	Telefone	string
	Email       string
}

func addStudent() error {
	var id string
	var name string
	var phone string
	var email string
	
	id = uuid.New().String()
	for true {
		fmt.Printf("Digite o nome do aluno novo: ")
		fmt.Scan(&name)
		if name != "" {
			break
		}
	}
	for true {
		fmt.Printf("Digite o telefone do aluno novo: ")
		fmt.Scan(&phone)
		if phone != "" {
			break
		}
	}
	for true {
		fmt.Printf("Digite o email do aluno novo: ")
		fmt.Scan(&email)
		if email != "" {
			break
		}
	}

	student := fmt.Sprint(id,",",name,",",phone,",",email,"\n")
	file,err := os.OpenFile("students.csv",  os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(student))
	if err != nil {
		return err
	}
	return nil
}

func loadStudents() ([]Student, string, error) {
	file, err :=os.ReadFile("students.csv")
	if err != nil {
		fmt.Println(err)
		return nil, "", err
	}
	content := string(file)
	allLines := strings.Split(content, "\n")
	var allStudents []Student
	
	for _, line := range allLines {
		words := strings.Split(line, ",")
		if len(words) == 4 {
			student := Student{
				Matricula: words[0],
				Nome: words[1],
				Telefone: words[2],
				Email: words[3],
			}
			allStudents = append(allStudents, student)
		}
	}
	return allStudents, content, nil
}

func listStudents(allStudents []Student) {
	fmt.Println("Lista de todos os alunos:")
	fmt.Printf("%3s | %36s | %10s | %10s | %28s \n", "ID","Matrícula", "Nome", "Telefone", "Email")
	for idx, s := range(allStudents) {
		fmt.Printf("%3d | %36s | %10s | %10s | %28s \n", idx + 1, s.Matricula, s.Nome, s.Telefone, s.Email)
	}
	fmt.Println()
}

func searchStudent(allStudents []Student) {
	var matricula string
	for true {
		fmt.Printf("Digite a matricula que você deseja procurar: ")
		fmt.Scan(&matricula)
		if matricula != "" {
			break
		}
	}
	fmt.Printf("%36s | %10s | %10s | %28s \n", "Matrícula", "Nome", "Telefone", "Email")
	for _,s := range(allStudents) {
		if s.Matricula == matricula {
			fmt.Printf("%36s | %10s | %10s | %28s \n", s.Matricula, s.Nome, s.Telefone, s.Email)
			break
		}
	}
	fmt.Println()
}

func saveContent(newContent string) error {
	file,err := os.OpenFile("students.csv",  os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	file.Write([]byte(newContent))

	return nil
}

func editStudent(allStudents []Student, file string) {
	var studentId int
	listStudents(allStudents)
	for true {
		fmt.Printf("Escolha o id do estudante que deseja alterar: ")
		fmt.Scan(&studentId)
		if studentId != 0 {
			break
		}
	}
	s := allStudents[studentId - 1]
	oldInfo := fmt.Sprint(s.Matricula,",",s.Nome,",",s.Telefone,",",s.Email)


	var name string
	var phone string
	var email string

	for true {
		fmt.Printf("Altere o nome do aluno: ")
		fmt.Scan(&name)
		if name != "" {
			break
		}
	}
	for true {
		fmt.Printf("Altere o telefone do aluno: ")
		fmt.Scan(&phone)
		if phone != "" {
			break
		}
	}
	for true {
		fmt.Printf("Altere o email do aluno: ")
		fmt.Scan(&email)
		if email != "" {
			break
		}
	}
	newInfo := fmt.Sprint(s.Matricula,",",name,",",phone,",",email)
	newContent := strings.Replace(file, oldInfo, newInfo, 1)
	saveContent(newContent)
}

func deleteStudent(allStudents []Student, content string) {
	var studentId int
	listStudents(allStudents)
	for true {
		fmt.Printf("Escolha o id do estudante que deseja deletar: ")
		fmt.Scan(&studentId)
		if studentId != 0 {
			break
		}
	}
	s := allStudents[studentId - 1]
	oldInfo := fmt.Sprint(s.Matricula,",",s.Nome,",",s.Telefone,",",s.Email, "\n")
	newContent := strings.Replace(content, oldInfo, "", -1)
	saveContent(newContent)
}


func main() {
	var choice int
	for true {
		allStudents, file, err := loadStudents()
		if err != nil {
			continue
		}
		fmt.Println("=== Menu de Opções ===")
		fmt.Println("1 - Incluir Aluno")
		fmt.Println("2 - Listar Alunos")
		fmt.Println("3 - Pesquisar Aluno por Matrícula")
		fmt.Println("4 - Alterar Aluno")
		fmt.Println("5 - Excluir Aluno")
		fmt.Println("6 - Sair")
		fmt.Printf("Escolha uma opção: ")
		fmt.Scan(&choice)
		if choice == 1 {
			addStudent()
		} else if choice == 2 {
			listStudents(allStudents)
		} else if choice == 3 {
			searchStudent(allStudents)
		} else if choice == 4 {
			editStudent(allStudents, file)
		} else if choice == 5 {
			deleteStudent(allStudents, file)
		} else if choice == 6 {
			fmt.Println("Saindo do programa...")
			break
		}
	}
}