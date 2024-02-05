// Pacote de execução
package main

//servidor
//import "net/http"

//Trata os valores da String e mostra os erros que surgirem
import (
	"fmt"
	"time"
)

/*
//Função soma
func soma(x int, y int) int {
	return x + y
}

//Função soma com return Booleano
func somaSit2(x int, y int) (int, bool) {
	if x > 10 {
		return x + y, true
	}
	return x + y, false
}*/

type NomeDoStruct struct { //1º letra Maiúscula Pública, Minúscula Privada
	nome  string
	email string
	idade int
}

type Empresas struct { //1º letra Maiúscula Pública, Minúscula Privada
	nome  string
	cnpj  int
	email string
	idade int
}

// Metodo da Empresas
func (e Empresas) getFullInfo() string { //aqui esta privada a 1º letra nome do metodo
	return fmt.Sprintf("Metodo --> Empresa: %s, CNPJ: %d, Email: %s, Idade: %d", e.nome, e.cnpj, e.email, e.idade)
}

// Metodo --> for/Sleep
func counter() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {

	//Call counter()
	fmt.Println("-----[Não Paralelo]-----")
	counter() //  não paralelo

	fmt.Println("-----[Paralelo]-----")
	go counter() // Thered's(T1) -> paralelo
	go counter() // Thered's(T2) -> paralelo
	counter()

	/*
		//parece construtor mas não é
		estrutura := NomeDoStruct{ //declarando os valores das variaveis com dois pontos(:)
			nome: "jackson", email: "j@email", idade: 31,
		}
		//Imprimindo os valores das variaveis
		fmt.Println("Nome:", estrutura.nome, "Email:", estrutura.email, "Idade:", estrutura.idade)
		//Adicionar novo valor
		estrutura.nome = "Jackson Neves"

		//Empresa
		dadosEmp := Empresas{"Jackson Business", 12345678919, "contact@jkbusiness.com", 25}
		// Print dos valores do metodo
		fmt.Println(dadosEmp.getFullInfo())

		//dadosEmp := Empresas{ nome: "Jackson Business", cnpj: 12345678919, email: "contact@jkbusiness.com", idade: 25, }
		fmt.Println("Empresa:", dadosEmp.nome, "CNPJ:", dadosEmp.cnpj, "Email:", dadosEmp.email, "Idade:", dadosEmp.idade)
	*/
}

/* func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

//EndPoint of API
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Dev"))
} */

/*
func main() {

	var a string //possui tipagem forte
	a = "Hello Dev"

	//aplica o tipo ao colocar :=
	b := "Aqui ele ja identifica o tipo de variavel"

	println(a)
	println(b)
	println("--------[SOMA]------")
	//Function Int
	println("Soma", soma(10, 55), "R$")
	println("--------[SOMA com Booleano]------")
	//Function Bool - Multiplo Valores
	resultado, status := somaSit2(10, 55) // aqui concatena(int e bool) adicionando os respectivos valores.
	println("Soma:", resultado, "is True or False?", status, "!!!")

} */
