package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	result, err := doSomething("input")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func doSomething(input string) (string, error) {
	if input == "" {
		return "", errors.New("input is empty")
	}
	var slice []int
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)
	// slice := []int{1, 2, 3}
	return "result", nil
}