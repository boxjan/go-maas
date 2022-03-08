package main

import (
	"encoding/json"
	"fmt"
	maas "github.com/boxjan/go-maas"
	"log"
)

func main() {
	client, err := maas.NewMaasClient(
		"http://127.0.0.1:5240/MAAS/",
		"2.0",
		"",
		0,
		"", nil)
	if err != nil {
		log.Fatalln(err)
	}

	ips, err := client.GetIpAddresses("all", "true")
	if err != nil {
		log.Fatalln(err)

	}
	for _, ip := range *ips {
		fabric, err := client.GetFabric(ip.Subnet.VLAN.FabricId)
		if err != nil {
			log.Fatalln(err)
		}
		t := map[string]string{}
		fmt.Printf("%s, %+v\n", json.Unmarshal([]byte(fabric.X["description"].(string)), &t), t)
	}
	fmt.Println(ips)
}
