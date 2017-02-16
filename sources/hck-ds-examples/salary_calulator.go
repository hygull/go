/*
https://www.hackerearth.com/practice/basic-programming/implementation/basics-of-implementation/practice-problems/algorithm/matrix-symmetry/

*/
package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	var salary float64
	var workingDays, selectedMonthNumber uint
	var year uint64 //Parse functions are of widest type
	var garbageConsumer string
	var err error

	/*Year calculation */
	yearStr := time.Now().String()[0:4]
	year, err = strconv.ParseUint(yearStr, 10, 64)
	year = 2016
	if err != nil {
		color.Red("[ERROR] Could not detect your System time properly...\n" +
			"It needs the datetime format like => 2017-01-06 21:49:17, So fix & retry")
		return
	}
	/*Year calculation ends*/

	monthsSlice := []string{"January", "February", "March", "April", "May",
		"June", "July", "August", "September", "October",
		"November", "December"}
	daysSlice := []uint{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	monthsToDaysMap := make(map[string]uint)

	choice := "y"
	for choice == "Y" || choice == "y" {

		//fmt.Println("\n*********************** Salary calculator ****************************\n")
		color.Blue("\n\n*********************** Salary calculator ****************************\n")
		//fmt.Println("Please select the month from the list below... \nfor which you want to calculate the salary...\nEnter 1/2/3/4 ... /12")
		color.Yellow("Please select the month from the list below... \nfor which you want to calculate the salary...\nEnter 1/2/3/4 ... /12")
		for index, month := range monthsSlice {
			monthsToDaysMap[month] = daysSlice[index]

			fmt.Println("[", index+1, "]\t", month)
		}

		//fmt.Print("\nFor which month you wanna give the salary?\n")
		color.Green("\nQue) For which month you wanna give the salary?\n")
		fmt.Println("--------------------------------------------------------------------")
		fmt.Print("Enter (Month number) : ")
		_, err = fmt.Scanf("%d", &selectedMonthNumber)

		fmt.Println("--------------------------------------------------------------------")

		if err != nil {
			color.Red("[ERROR] You entererd an invalid integer, please enter any number ranged from  1...12")
			//fmt.Println("[ERROR] You entererd an invalid integer, please enter any number ranged from  1...12")
			fmt.Scanf("%s", &garbageConsumer)
			continue
		}
		if !(selectedMonthNumber > 0 && selectedMonthNumber < 13) {
			//fmt.Println("[ERROR] You entered a wrong month number,please enter any number ranged from  1...12")
			color.Red("[ERROR] You entered a wrong month number,please enter any number ranged from  1...12")
			continue
		}

		noOfDays := daysSlice[selectedMonthNumber-1] //noOfDays=1
		if selectedMonthNumber == 2 {
			if year%400 == 0 {
				noOfDays = 29
			} else {
				if year%100 == 0 {
					noOfDays = 28
				} else {
					if year%4 == 0 {
						noOfDays = 29
					}
				}
			}
		}

		//fmt.Println("You selected       ==> ", monthsSlice[selectedMonthNumber-1])
		color.Cyan("You selected       ==> %s", monthsSlice[selectedMonthNumber-1])
		//fmt.Println("Numeber of days    ==> ", daysSlice[selectedMonthNumber-1])
		color.Cyan("Numeber of days    ==> %d", noOfDays)

		//fmt.Print("\nWhat is the monthly salary of your employee?\n")
		color.Green("\nQue) What is the monthly salary of your employee?\n")
		fmt.Println("--------------------------------------------------------------------")
		fmt.Print("Enter (Salary)       : ")
		_, err = fmt.Scanf("%f", &salary)
		fmt.Println("--------------------------------------------------------------------")

		if err != nil {
			//fmt.Println("[ERROR] You entererd an invalid integer, please enter your valid salary")
			color.Red("[ERROR] You entererd an invalid integer, please enter your valid salary")
			fmt.Scanf("%s", &garbageConsumer)
			continue
		}
		if !(salary > 0.0) {
			//fmt.Println("[ERROR] Salary should be greater than 0.")
			color.Red("[ERROR] Salary should be greater than 0.")
			continue
		}

		//fmt.Print("\nFor how many days your employeed worked for this month?\n")
		color.Green("\nQue) For how many days your employeed worked for this month?\n")
		fmt.Println("--------------------------------------------------------------------")
		fmt.Print("Enter (Working days) : ")
		fmt.Scanf("%d", &workingDays)
		fmt.Println("--------------------------------------------------------------------")

		if err != nil {
			//fmt.Println("[ERROR] You entererd an invalid integer, please enter your valid salary")
			color.Red("[ERROR] You entererd an invalid integer, please enter your valid salary")
			fmt.Scanf("%s", &garbageConsumer)
			continue
		}
		if !(workingDays >= 0) {
			//fmt.Println("[ERROR] Working days cannot be negative.")
			color.Red("[ERROR] Working days cannot be negative.")
			continue
		}
		// fmt.Println(salary)
		// fmt.Println(salary / float64(workingDays))
		totalSalary := float64(workingDays) * (salary / float64(noOfDays))

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		fmt.Println("*************** Salary details ***********************")
		fmt.Println("Month                   : ", monthsSlice[selectedMonthNumber-1])
		fmt.Println("Working days            : ", workingDays)
		fmt.Println("Monthly salary          : ", salary)
		fmt.Printf("Salary(for this month)  :  %.2f", totalSalary)

		fmt.Print("\n\nDo you want to continue...<Y/N> : ")
		//color.Magenta("\n\nDo you want to continue...<Y/N> : ")
		fmt.Scanf("%s", &choice)
	}
	//fmt.Println("\nBye...\n")

	fmt.Println("--------------------------------------------------------------------")
	color.Magenta("\nBye...")
}
