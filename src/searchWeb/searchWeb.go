package searchWeb

import (
	"fmt"
	"net/http"
)

func GetUrl(url string) *http.Response { // Se a primeira letra for maiúscula, então a função é visível fora do pacote.
	res, err := http.Get(url)
	if err != nil {
		panic("Error")
	}

	fmt.Println(res.Body)
	return res
}
