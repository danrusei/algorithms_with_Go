package main

import (
	"github.com/danrusei/algorithms_with_go/other_ds/hashtable"
)

func main() {

	hashmap := hashtable.New()
	hashmap.AddKeyValue("first", 1)
	hashmap.AddKeyValue("second", 2)
	hashmap.AddKeyValue("thrid", 3)
	hashmap.AddKeyValue("fourth", 4)
	hashmap.AddKeyValue("fifth", 5)
	hashmap.AddKeyValue("sixth", 6)
	hashmap.AddKeyValue("seventh", 7)
	hashmap.AddKeyValue("eigth", 8)
	hashmap.AddKeyValue("nineth", 9)
	hashmap.AddKeyValue("tenth", 10)
	hashmap.AddKeyValue("eleventh", 11)

	hashmap.String()
}
