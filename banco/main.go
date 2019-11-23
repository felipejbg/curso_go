package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	conta1 := contas.ContaCorrente{Titular: "Silvia", Saldo: 500}
	contaA := contas.ContaCorrente{Titular: "Joao", Saldo: 300}
	contaB := contas.ContaCorrente{Titular: "Maria", Saldo: 100}

	fmt.Println(contaA)
	fmt.Println(contaB)

	fmt.Println(contaA.Transferir(150, &contaB))

	fmt.Println(contaA)
	fmt.Println(contaB)

	fmt.Println(conta1.Saldo)
	fmt.Println(conta1.Sacar(200))
	fmt.Println(conta1.Saldo)
	fmt.Println(conta1.Sacar(400))
	fmt.Println(conta1.Depositar(1000))
	fmt.Println(conta1.Saldo)
	fmt.Println(conta1.Sacar(-100))
}
