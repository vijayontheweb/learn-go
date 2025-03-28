package main

func main() {
	value := "initial"
	println("initial value:", value)
	changeValue(value)
	println("value after changeValue:", value)
	changeValueByPointer(&value)
	println("value after changeValueByPointer:", value)
}

func changeValue(value string) {
	value = "changed"
}

func changeValueByPointer(value *string) {
	*value = "changed"
}
// Output:
// initial value: initial
// value after changeValue: initial
// value after changeValueByPointer: changed
