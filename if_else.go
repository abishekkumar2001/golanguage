package main

import "fmt"

func main(){

	//If
	var n int
	fmt.Println("Enter a number : ")
	fmt.Scan(&n)
	if n<0 {
		fmt.Println("The number is Negative")
	}

	//If ... Else
	var a int
	fmt.Println("Enter a number : ")
	fmt.Scan(&a)
	if n%2==0{
		fmt.Println("The number is Even")
	}else{
		fmt.Println("The number is odd")
	}

	//If ... Elseif ... else
	var n1,n2,n3 int
	fmt.Println("Enter the first number : ")
	fmt.Scanln(&n1)
	fmt.Println("Enter the second number : ")
	fmt.Scanln(&n2)
	fmt.Println("Enter the third number : ")
	fmt.Scanln(&n3)

	if n1>n2 && n1>n3{
		fmt.Printf("%d is the greatest of three numbers",n1)
	}else if n2>n1 && n2>n3{
		fmt.Printf("%d is the greatest of three numbers",n2)
	}else{
		fmt.Printf("%d is the greatest of three numbers",n3)
	}

	//Nested if
	var b int
	fmt.Println("Enter the value of b :")
	fmt.Scanln(&b)
	if b>0 {
		if b%2==0 {
			fmt.Println("The number is positive and even")
		}
	}
	
	//Loops in go - for loop
	var i,limit,sum int
	sum=0
	fmt.Println("Enter the limit : ")
	fmt.Scanln(&limit)
	for i=0; i<limit; i++{
		sum=sum+i;
	}
	fmt.Println("The sum is ",sum)

	//Getting input from user for two dimensional array using for loop
	var rows,columns int
	var j int = 0
	fmt.Println("Enter the number of rows : ")
	fmt.Scanf("%d",&rows)
	fmt.Println("Enter the number of columns");
	fmt.Scanf("%d",&columns)
	var twoArray = make([][]int,rows) 
	for i := range twoArray{ //creating arrays for rows
		twoArray[i]=make([]int,columns)
	}
	for i=0; i<rows; i++{
		for j=0; j<rows; j++{
			fmt.Scanf("%d",&twoArray[i][j])
		}
	}
	fmt.Println(len(twoArray))

	for i=0; i<rows; i++{
		for j=0; j<rows; j++{
			fmt.Printf("%d",twoArray[i][j])
		}
		fmt.Printf("\n")
	}

	//Continue statement
	fmt.Println("Enter the limit : ")
	fmt.Scanln(&limit)
	for i=0; i<=limit; i++{
		if i==3 {
			continue
		}
		fmt.Println(i)
	}

	//Break statement
	fmt.Println("Enter the limit : ")
	fmt.Scanln(&limit)
	for i=0; i<=limit; i++{
		if i==7 {
			break
		}
		fmt.Println(i)
	}

	//nested loops
	var number int =5
	for i=0; i<number; i++{
		fmt.Printf("At Iteration %v \n\n",i+1)
		for j=0; j<number;j++{
			fmt.Printf("The value of i is %d and j is %d \n",i,j)
		}
		fmt.Println("")
	}

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