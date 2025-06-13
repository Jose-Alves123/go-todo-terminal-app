// Package services implements utilities to support the main application,
// being with the use of os helper function or ainding with the arriging of
// data in the database
package services

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Allows user to input information about the data of a todo
//
// Data to be scanned includes title (string), description (string),
// state (uint8)
func DataToCreateToDo() (string, string, uint8, error) {
	fmt.Println("Por favor insira os dados de um novo to do")
	fmt.Print("Título: ")
	title, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		return "", "", 0, err
	}
	title = strings.TrimSuffix(title, "\n")

	fmt.Print("Descrição: ")
	description, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		return "", "", 0, err
	}
	description = strings.TrimSuffix(description, "\n")

	fmt.Print("Estado (0 -> ToDo, 1 -> Doing, 2 -> Done): ")
	state, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		return "", "", 0, err
	}
	state = strings.TrimSuffix(state, "\n")

	stateInt, err := strconv.ParseUint(state, 10, 8)
	if err != nil {
		return "", "", 0, err
	}

	return title, description, uint8(stateInt), nil

}
