package hashtable

import (
	"fmt"
	"math"

	"github.com/danrusei/algorithms_with_go/linkedlist"
)

// kv stores the HashMap key-value pair
type kv struct {
	key   string
	value interface{}
}

// the linked list that stores the chained value of each Bucket
//type llist struct {
//	linkedlist.Linkedlist
//}

// number the buckets, it is in fact the length of the array, which is the underlying structure of the hashtable
const bucketCount = 16

// HashTable data structure is the central structure of the package
type HashTable struct {
	data [bucketCount]*linkedlist.Linkedlist
}

// New is the HashTable constructor
func New() *HashTable {
	return &HashTable{
		[bucketCount]*linkedlist.Linkedlist{},
	}
}

// the hash algorithm to convert the key to an integer of 0 - bucketCount as the index of the array to store/retrieve value.
// the hash algorithm should be as random as possible
// so that the newly added key-value pairs are evenly distributed under each array.
// Java’s implementation of hash functions for strings is a good example:
// key[0]·31n-1 + key[1]·31n-2 + … + key[n-1]  -- where n is the length of the key
func hash(key string) int {
	hash := 0
	for index, char := range key {
		hash += int(char) * int(math.Pow(31, float64(len(key)-index+1)))
	}

	//the reminder is the selected bucket
	bucket := hash % bucketCount
	return bucket
}

// AddKeyValue pairs to HashTable
func (myMap *HashTable) AddKeyValue(key string, value interface{}) {
	bucket := hash(key)

	if myMap.data[bucket] == nil {
		myMap.data[bucket] = linkedlist.NewLinkedlist()
		//AddAtBeg adds a Node at the beginning of the linked list
		myMap.data[bucket].AddAtBeg(kv{key, value})
		return
	}
	myMap.data[bucket].AddAtBeg(kv{key, value})
}

// GetValueForKey retrieves the value for a gven key
func (myMap *HashTable) GetValueForKey(key string) (interface{}, bool) {
	bucket := hash(key)

	if myMap.data[bucket] == nil {
		return 0, false
	}

	//check if the key == first node of the linked list
	// assertion is required as the Node.Val == interface{}
	cur := myMap.data[bucket].Head
	if cur.Val.(kv).key == key {
		return cur.Val.(kv).value, true

	}
	// if not, traverse the linked list, untill the value is found
	for ; cur.Next != nil; cur = cur.Next {
		node := cur.Val.(kv)
		if node.key == key {
			return node.value.(int), true
		}

	}
	// else return false
	return 0, false
}

// utility function to print the HashTable
func (myMap *HashTable) String() {
	for i := range myMap.data {
		if myMap.data[i] == nil {
			fmt.Printf("Bucket %d has no Nodes\n", i)
			continue
		}
		fmt.Printf("Bucket %d has the following Nodes:", i)
		fmt.Println(myMap.data[i])
	}
}
