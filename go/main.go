package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	domain, newNameservers := parseParams()
	Remap(context.TODO(), domain, newNameservers)
}

func parseParams() (string, []string) {
	domain := os.Args[1]
	fmt.Printf("domain: %s\n", domain)
	newNameservers := os.Args[2:]
	return domain, newNameservers
}
