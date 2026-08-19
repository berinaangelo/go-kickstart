package main

import "fmt"

func main() {
	// initArray()
	// initMaps()
	// initStruct()

	ll := &LinkedList{}
	ll.Create(55)
	ll.Create(68)
	ll.Create(87)
	ll.Create(29)

	ll.Display()

	data, _ := ll.Show(1)
	fmt.Println(data.data) // index accepted as param

	fmt.Println(ll.Search(29)) // data accepted as param

	ll.Update(0, 65)
	ll.Display()

	ll.Delete(0)
	ll.Display()
}
