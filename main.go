package main
import "fmt"
func main(){
	fmt.Println("Hello, world!");
	if true{
		fmt.Println("this is true");
	} else {
		fmt.Println("this is false");
	}
	for i:=0;i<5;i++{
		fmt.Println(i);
	}
	type Person struct {
		name string
		age int
	}
	var p Person
	p.name = "John"
	p.age = 30
	fmt.Println(p)

	//callback function
	func(callback func(string)){
		callback("Hello from callback");
	}(func(message string){
		fmt.Println(message);
	})
	//differ function with lifo principle
	defer fmt.Println("This will be printed last");
	defer fmt.Println("This will be printed second");
	defer fmt.Println("This will be printed third");
	fmt.Println("This will be printed first");	

}