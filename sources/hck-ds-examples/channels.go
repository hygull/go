/*
	{
		"cretaed_after" : "Mon Dec 12 22:35:30 IST 2016"
		"aim_of_program": "channels (Printing table of numbers from 1 to 10)"
		"coded_by" 	: "Rishikesh Agrawani"
	}
*/
package main

import "fmt"

func main() {
	/* Defining a channel for storing integers */
	intsChan := make(chan int)

	/*Defining go routine*/
	go func() {
		for i := 1; i < 11; i++ {
			intsChan <- i //Sending numbers from 1-10 into channel
		}
	}()

	for i := 1; i < 11; i++ {
		n := <-intsChan //Receiving numbers from channel one by one in each iteration of for loop and assigning it to variable n
		for j := 1; j < 11; j++ {
			fmt.Print(n*j, "\t")
		}
		fmt.Println()
	}

	fmt.Println("")
}

/*
1	2	3	4	5	6	7	8	9	10
2	4	6	8	10	12	14	16	18	20
3	6	9	12	15	18	21	24	27	30
4	8	12	16	20	24	28	32	36	40
5	10	15	20	25	30	35	40	45	50
6	12	18	24	30	36	42	48	54	60
7	14	21	28	35	42	49	56	63	70
8	16	24	32	40	48	56	64	72	80
9	18	27	36	45	54	63	72	81	90
10	20	30	40	50	60	70	80	90	100
*/
