package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3}
	fmt.Println(len(slice1)) //To find the length of slice
	fmt.Println(cap(slice1)) //To find the maximum capacity the slice can grow and shrink
	fmt.Println(slice1)

	slice2 := []string{"This","is","Go","Program"}
	fmt.Println(len(slice2)) //To find the length of slice
	fmt.Println(cap(slice2)) //To find the maximum capacity the slice can grow and shrink
	fmt.Println(slice2)
	
	//Slicing from an array
	var arr = [...]int{10,25,15,33,29}
	slice := arr[1:3] //slicing from an existing array
	fmt.Println(slice)
	fmt.Println(len(slice)) //Length of slice

	//Creating slice from make function
	slice3 := make([]int,3,7) //SYNTAX :sliceName := make([]datatype,length,capacity)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))	

	//Assigning values to slice3 using loop
	var i int
	fmt.Println("Enter the elements of slice3") // We can also assign values in this way slice3[0]=1 slice3[1]=2 slice3[2]=3
	for i=0; i<len(slice3); i++{
		fmt.Scan(&slice3[i])
	}
	fmt.Println("The elements of slice 3 are")
	for i=0; i<len(slice3); i++{
		fmt.Print(slice3[i],"\t")
	}

	fmt.Print("\n")

	fmt.Println("The elements of slice3 are",slice3)

	//Accessing elements of slice
	fmt.Println(slice1[1])

	//Changing the elements of a slice
	slice1[1]=23
	slice2[3]="Example"
	fmt.Println(slice1)
	fmt.Println(slice2)

	//Appending elements to the slice

	//Appending elements at the end
	slice4 := []int{1,2,3,4}
	fmt.Println("Length and Capacity before appending")
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	slice4 = append(slice4,20,20)
	fmt.Println(slice4)
	fmt.Println("Length and Capacity after appending")
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))


	slice5 := []int{1,2,3,4,5,6}
	fmt.Println("Length and Capacity before appending")
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))
	slice5 = append(slice5,7,8,9,10,11,12,13)
	fmt.Println("Length and Capacity after appending")
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))
	
	//Appending one slice to another slice
	slice6 := append(slice3,slice5...) //SYNTAX slice1 := append(slice2,slice3)
	fmt.Println(slice6)

	slice7 := append(slice4,slice5...) 
	fmt.Println(slice7)

	//Changing the length of a slice
	var array = []int{1,2,3,4,5,6,7}
	sliceArray1 := array[1:5]
	fmt.Println(len(sliceArray1))
	sliceArray1 = array[1:3]
	fmt.Println(len(sliceArray1))

	//Copy function
	sliceArray2 :=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	n1 := sliceArray2[:len(sliceArray2)-10]
	n2 := make([]int,len(n1))
	copy(n2,n1)
	fmt.Println(n2)

	//Range in Go 
	var fruits = []string{"Apple","Orange","Banana","Guava","Grapes"}
	for idx, val := range fruits{
		fmt.Printf("%v %v\n",idx,val)
	}

	//Range without pinting index value 
	for _, val := range fruits{
		fmt.Printf("%v\n",val)
	}
}
