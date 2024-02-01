package math

import "testing"

//The testing file name must ends with _test.go.
//Go testing files are always located in the same folder, or package, where the code they are testing resides.

func TestAdd(t *testing.T) { //The function signature must be like this only.
	//Function name should be Test followed by a suffix with its first letter capital.
	//The first and only parameter must be t *testing.T.

	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	received := Subtract(6,4)
	ans :=2
	
	if received != ans {
		t.Errorf("Received %d, ans %d",received,ans)
	}

	//To run testing  1. go mod init example.com/<package name> here testing is the package name 2.go mod tidy 3. go test

}

