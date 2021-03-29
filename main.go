package main

import "github.com/svrakata/ekatte_api/ekatte"

func main() {
	err := ekatte.ExtractData()

	if err != nil {
		panic(err)
	}

	err = ekatte.ParseEkatte()

	if err != nil {
		panic(err)
	}

}
