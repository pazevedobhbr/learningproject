package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(3200))

	//com sinal ...int8  int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("o valor máximo dom int é", i1)
	fmt.Println("O VALOR DO TIPO I1 É", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mmapeameno da tabela Unicode (inte32)
	fmt.Println("o topo de i2", reflect.TypeOf(i2))
	fmt.Println(i2)

	//números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("o tipo de x é", reflect.TypeOf(x))
	fmt.Println("o tipo do valor literal de 49.99 é", reflect.TypeOf(49.99))
	fmt.Println(x)

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println((!bo))

	//string
	s1 := "Ola meu nome é Paulo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	//string com multiplas linhas (usa crase)

	s2 := `Olá
			meu 
			nome 
			é 
			Paulo`
	fmt.Println(s2)
	fmt.Println("O tamnho da string é", len(s2))

	//char
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println("O valor  de char é", char)

}
