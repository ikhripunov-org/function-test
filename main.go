package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var array []interface{}
	if err := json.Unmarshal([]byte(os.Args[1]), &array); err != nil {
		fmt.Printf("Error while decoding %v\n", err)
	}
	integer, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Error while decoding %v\n", err)
	}
	fmt.Println(count(array, integer))
}

func count(array []interface{}, integer int) int {
	var counter = 0
	for _, e := range flattenAndCast(nil, array) {
		if e == integer {
			counter = counter + 1
		}
	}
	return counter
}

// Does 2 things since json.Unmarshal only produces float64 for interface value
func flattenAndCast(args []int, arrayOfArrays interface{}) []int {
	if array, ok := arrayOfArrays.([]interface{}); ok {
		for _, nestedArray := range array {
			args = flattenAndCast(args, nestedArray)
		}
	} else {
		args = append(args, int(arrayOfArrays.(float64)))
	}
	return args
}
