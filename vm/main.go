package main

func main() {
	// create a new vm with the code
	vm := NewVM([]Code{NewPushCode(1), NewPushCode(2), NewPlusCode(), NewPushCode(3), NewMultiplyCode()})
	// run the vm
	result := vm.Run()
	// print the result
	showVMResult(result)
}
