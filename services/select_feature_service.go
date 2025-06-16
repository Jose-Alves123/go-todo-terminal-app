package services

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Asks the user and scans to obtain the option desrided from the user.
func SelectFeature() (int, error) {
	fmt.Print("Por favor indique o que pretende.\n0 - Listar ToDos\n1 - Criar ToDo\n2 - Eliminar dados de BD\nR: ")
	answer, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		return 0, err
	}

	answer = strings.TrimSuffix(answer, "\n")
	number, err := strconv.Atoi(answer)
	if err != nil {
		return 0, err
	}

	if number < 0 || number > 2 {
		return 0, errors.ErrUnsupported
	}

	return number, nil
}
