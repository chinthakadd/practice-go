package main

import (
	"practice-go/util"
)

//It is illegal in Go to import a package and not to use it anywhere in the code. The compiler will complain if you do so.
//The reason for this is to avoid bloating of unused packages which will significantly increase the compilation time.

// Error Silencer is a mechanism to allow importing packages that are not used yet while doing development.
// it helps test out package level init() method logics etc.
var _ = util.ReturnSomething() //error silencer

func main() {

}
