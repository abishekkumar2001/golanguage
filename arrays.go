package main

import "fmt"

func main(){

	//Array declaration using var keyword and length
	var arr1 = [5]int{5,10,15,20,25}
	//Array declaration using var keyword without length 
	var arr2 = []string{"Apple","Orange","Grapes"}
	//Array declaration using := sign and length
	arr3 := [4]float32{1.1,2.2,3.3,4.4}
	//Array declaration using := sign without length
	arr4 := []int{11,21,31,41,51}
	fmt.Println(arr1) 
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	//Printing the elements of the array using println method
	fmt.Println("The elements of the first array are")
	var i int
	for i=0; i<len(arr1); i++{
		fmt.Println(arr1[i])
	}

	//Printing the elements of the array using printf function
	fmt.Println("The elements of the second array are")
	for i:=0; i<len(arr4); i++{
		fmt.Printf("%d",arr4[i]) //Printf and scanf have similar syntax as in c
	}

	//Accessing the elements of the array
	fmt.Println(arr1[2]) // It will print the element at index 2 of arr1
	fmt.Println(arr2[0]) // It will print the element at index 0 of arr2
	fmt.Println(arr3[3]) // It will print the element at index 3 of arr3

	//Changing the value of the array
	arr1[4]=250
	arr2[1]="Mango"
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	//Array Initialization
	var array1=[7]int{} //not initialized
	var array2=[7]int{1,2,3} //partially initialized
	var array3=[7]int{1,2,3,4,5,6,7} //Completely initialized
	var array4=[3]string{} 
	array5:=[5]string{"One","Two","Three","Four","Five"}
	var array6=[5]string{"One","Two","Three"}
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)
	fmt.Println(array5)
	fmt.Println(array6)

	//Initializing only specific elements
	var a1=[5]int{0:50,2:150}
	a2 :=[5]string{2:"Green",4:"White"}
	fmt.Println(a1)
	fmt.Println(a2)

	//Finding length of an array
	var car=[]string{"BMW","Mahindra","Fird"} 
	marks:=[]int{60,70,80,90,100}
	fmt.Println(len(car))
	fmt.Println(len(marks))

	//Getting array as input from user
	var n int;
	fmt.Println("Enter the size of the array")
	fmt.Scanln(&n)
	
	a3 := make([]int,n)

	fmt.Println("Enter the elements of the array...")
	for i=0; i<n; i++{
		fmt.Printf("Enter the element at index %d ",i)
		fmt.Scanf("%d",&a3[i])
	}

	fmt.Println("The Elements of the array are...")

	for i=0; i<n; i++{
		fmt.Printf("%d\t",a3[i])
	}

	fmt.Println("a3 is ",a3)
}
