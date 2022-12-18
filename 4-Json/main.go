package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {

	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(conta)
	//também possível declarar em uma só linha, pois apenas usamos o método Encode de encoder, sendo possível chamar o Encode diretamente:
	//err = json.NewEncoder(os.Stdout).Encode(conta)

	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"n":2, "s":200}`)
	var contaX Conta

	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Saldo, contaX.Numero)

}
