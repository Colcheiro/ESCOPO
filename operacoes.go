package main
 
import "fmt"
type Jogador struct{
	nome string
	vida int
	nivel int

func main (){
j1:= Jogador {"Logan",100,3}
fmt.Println(j1.nome)
}