package main

func main() {
	err := ExtractData()

	if err != nil {
		panic(err)
	}

	err = ParseEkatte()

	if err != nil {
		panic(err)
	}

}
