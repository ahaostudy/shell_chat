package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLine(s string) string {
	scanner := bufio.NewScanner(os.Stdin)
	if len(s) > 0 {
		fmt.Print(s)
	}
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}
