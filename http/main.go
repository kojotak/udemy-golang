package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// chci vypsat obsah webove stranky

	//1. pristup
	//bs := make([]byte, 99999) //make = vezme dany typ a pripravi ho na danou velikost
	//resp.Body.Read(bs)        //nacti svuj obsah do bs
	//fmt.Println(string(bs))   //musim pretypovat byte slice na string

	//2. zkraceny pristup
	//io.Copy(os.Stdout, resp.Body) //zkopiruj do stdout to co je v body - vyuziva dvaou rozhrani: writer a reader

	//3.pristup - chci pouzit vlastni writer
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil //nil = bez chyby, ale musime vratit pocet zapsanych bajtu
}
