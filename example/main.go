package main

import (
	"fmt"
	maas "github.com/boxjan/go-maas"
)

func main() {
	client, err := maas.NewMaasClient(
		"http://127.0.0.1:5240/MAAS/",
		"2.0",
		"",
		0,
		"", nil)
	fmt.Println(client, err)
}
