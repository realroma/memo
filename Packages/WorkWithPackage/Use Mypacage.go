package main

import (
	"Project/memo/Packages/WorkWithPackage/OtherTry"
	"Project/memo/Packages/WorkWithPackage/mypackage"
	"fmt"
)

func main() {
	fmt.Println(mypackage.Do()) //Print "1"
	fmt.Println(OtherTry.Do())  // Print "1 /n 1"
}
