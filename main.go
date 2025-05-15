package main
 
import "fmt"
type Jogador struct{
	nome string
	vida int
	nivel int
}
func exibeDados (j Jogador){
	fmt.Println("O nome do jogador é", j.nome)
	fmt.Println("A vida do jogador é", j.vida)
	fmt.Println("O nível do jogador é", j.nivel)
}

func main (){
j1:= Jogador {"Logan",100,3}
exibeDados(j1)
}
