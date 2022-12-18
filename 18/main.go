package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	tam, err := f.Write([]byte("full check sum"))
	//tam, err := f.WriteString("check")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Success! Size: %d bytes\n", tam)
	f.Close()
	//leitura do arquivo

	arquivo, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	println(string(arquivo))

	//leitura de pouco em pouco no arquivo
	arquivo2, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
