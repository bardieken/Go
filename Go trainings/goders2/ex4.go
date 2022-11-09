package main

import "fmt"

func main() {
	//map
	baskentler := map[string]string{"Turkey": "Istanbul", "france": "paris"}
	for key := range baskentler {
		fmt.Println("Map item: Capital of", key, "is", baskentler[key])
	}
}
