package main
import "fmt"
func main(){
	commandKeySymbol:="\u2318";
	fmt.Println("Command key symbol on MAC OS X : ", commandKeySymbol);
	fmt.Printf("%T\n",commandKeySymbol);
	
	c:='a';
	appleSymbol:='';
	fmt.Printf("%+q\n%+q , %v, %c\n",c,appleSymbol,appleSymbol,appleSymbol);
}

