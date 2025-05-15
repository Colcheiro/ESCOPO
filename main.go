package main
 
import "fmt"


type Carro struct{
	modelo string
	ano int
	velocidade int
}

func estruturaCarro (c Carro){
	fmt.Println("O modelo do carro é", c.modelo)
	fmt.Println("O ano do carro é", c.ano)
	fmt.Println("A velocidade do carro é", c.velocidade)
}

func main(){
	c1:= Carro {"Ford Mustang",2022, 250}
	estruturaCarro(c1)
}
