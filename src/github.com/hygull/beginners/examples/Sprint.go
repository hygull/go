/*
 @ Date of creation : 06/09/2016.
 @ Aim of program   : To use Sprint in Go.
 @ Code by          : Rishikesh Agrawani.
 */
package main
import "fmt"

func main(){

fmt.Println(fmt.Sprint("This"," is ","great."));
fmt.Print("\n",hello(),"\n"); 
}

func hello() string {
  return fmt.Sprint("This is ","Rishikesh");
}
