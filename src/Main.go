package main

import (
	"ZzEmu"
	"fmt"
)

// type Animal struct {
// 	nome string
// 	cor  string
// }

// type Generico interface {
// 	andar() string
// 	comer() string
// }

// func (a *Animal) andar() string {
// 	return "O " + a.nome + " trotar"
// }

// func (a *Animal) comer() string {
// 	return "O " + a.nome + " come feno"
// }

func main() {

	par := ZzEmu.RegisterPair{}

	fmt.Println("Valor " + fmt.Sprint(par.Get()))

	// var pA *Animal = new(Animal)
	// pA.nome = "corvo"
	// pA.cor = "preto"

	// var pAnimal Generico = pA //new(Animal{nome: "aaa"}) //&Animal{nome: "onca", cor: "branca"} //new(Animal)
	// //pAnimal.nome = "burro"
	// //pAnimal.cor = "preto"

	// //cavalo Animal = Animal{nome: "cavalo", cor: "preta"}

	// fmt.Println(pAnimal.andar() + pAnimal.comer())

}
