package utils

import (
	"bufio"
	"os"
)

type AccountManager struct {
	AccountList []string
}

func (am *AccountManager) LoadFromFile(filename string) (int, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		am.AccountList = append(am.AccountList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return len(am.AccountList), nil
}
