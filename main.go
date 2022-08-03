package main

import "fmt"

func main() {
	inConf, _ := InfaktConfigFromJSON("config.json")
	hvConf, _ := HarvestConfigFromJSON("config.json")
	fmt.Printf("%+v\n", inConf)
	fmt.Printf("%+v\n", hvConf)
}
