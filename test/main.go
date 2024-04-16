package main

// Hello World表示する用
import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	result, err := doSomething("input")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
	// http.HandleFunc("/api", onigiri.ApiHandler)
	// http.ListenAndServe(":8080", nil)
}

func doSomething(input string) (string, error) {
	if input == "" {
		return "", errors.New("input is empty")
		//errorsパッケージのNewという関数を使用
	}
	return "result", nil
}