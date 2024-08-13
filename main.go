package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type BrasilAPIJson struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type ViaCepJson struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func getAdrress(ch chan interface{}, url string) {

	client := http.DefaultClient

	req, err := client.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer request a URL: ", url)
		fmt.Println(err)
		return
	}
	defer req.Body.Close()

	resp, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Erro ao ler corpo da resposta: ", resp)
		fmt.Println(err)
		return
	}

	var result BrasilAPIJson
	err = json.Unmarshal(resp, &result)
	if err != nil {
		fmt.Println("Erro unmarshaling json.")
	}

	ch <- result

}

func main() {

	c1 := make(chan interface{})
	c2 := make(chan interface{})
	url1 := "https://brasilapi.com.br/api/cep/v1/01153000"
	url2 := "http://viacep.com.br/ws/50670350/json/"

	go getAdrress(c2, url2)
	go getAdrress(c1, url1)

	select {
	case result := <-c1:
		fmt.Println("chegou", result)

	case result := <-c2:
		fmt.Println("viacep", result)

	case <-time.After(time.Second):
		fmt.Println("Timeout Error :D")

	}

}
