package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ProceedWithInfaktInvoice(ms MonthlySummary) bool {
	if len(ms) == 0 {
		return false
	}
	fmt.Println("Here are the hours fetched from Harvest system:")
	fmt.Println()
	ms.Print()
	fmt.Println()
	fmt.Println("Should I continue and create an invoice draft in Infakt? (y/n) ")

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		if strings.ToLower(sc.Text()) == "y" {
			return true
		} else {
			return false
		}

	}
	return false
}
