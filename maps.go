package main

import "fmt"

func main(){

	var map1 = map[string]int{"Tamil":95,"English":88,"Mathematics":94,"Science":94,"Social Science":98}

	//SYNTAX : var mapName = map[key type]value type{key1 : value1, key2 : value2, .... , keyn : valuen} 
	//OR b := map[key type]value type {key1 : value1, key2 : value2, .... , keyn : valuen}

	map2 := map[string]int{"Rohit Sharma":45,"Fazel":1,"Abinesh":4,"Rossi":46}

	fmt.Println(map1)	
	fmt.Println(map2)

	//Creating maps using make function

	var map3 = make(map[string]string) //Map3 is empty now
	map3["Language1"] = "Tamil"
	map3["Language2"] = "English"
	map3["Language3"] = "Malayalam"

	//Map 3 is having values now

	fmt.Println("The contents of map 3 are ...")
	fmt.Println(map3)

	//Adding and Changing values of map

	map1["Science"] = 92
	map3["Language4"] = "Telugu"
	fmt.Println(map1)

	//Accessing values of map

	var n1 = map1["Social Science"]
	n2 := map3["Language3"]
	fmt.Printf("The value of n1 is %d and the value of n2 is %s \n",n1,n2)

	//Deleting the elements of map

	map4 := map[string]int{"Apple":1, "Orange":2, "Banana":3, "Grapes":4}
	delete(map4,"Banana")
	fmt.Println(map4)
	keysToDelete := []string{"Orange","Grapes"}
	for _, key := range keysToDelete{
		delete(map4,key) //Here key is the value of keysToDelete.
	}
	fmt.Println(map4)

	//Getting input from the user

	var map5 map[string]int //Declaration
	var key string
	var val int
	var n int 
	fmt.Println("Enter the value of n : ")
	fmt.Scanln(&n)
	
	map5 = make(map[string]int) //Creating an empty map using make function

	for i:=0; i<n; i++ {
		fmt.Printf("Enter the key value : ")
		fmt.Scanln(&key)
		fmt.Println("Enter the value : ")
		fmt.Scanln(&val)

		//Storing key and val in map

		map5[key]=val
	}

	fmt.Println("The contents of map5 are ... ")
	fmt.Println(map5)
}













