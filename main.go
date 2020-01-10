package main

import (
	"fmt"
	"strings"
)

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
func main() {
	var data = []string{"atrisa", "endya", "nur", "hidayah"}
	var dataContainsi = filter(data, func(each string) bool {
		return strings.Contains(each, "i")
	})
	var dataLenght7 = filter(data, func(each string) bool {
		return len(each) == 7
	})

	fmt.Println("data asli \t\t:", data)
	// data asli : [atrisa endya nur hidayah]

	fmt.Println("filter ada huruf \"i\"\t:", dataContainsi)
	// filter ada huruf "o" : [atrisa] [hidayah]

	fmt.Println("filter jumlah huruf \"7\"\t:", dataLenght7)
	//filter jumlah huruf "6" : [hidayah]
}
