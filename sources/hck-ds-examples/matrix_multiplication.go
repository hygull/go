package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
)

func main() {
	var selectedMatrixCol int
	var option, optionPrev int
	var m1, m2 [][]uint                  //Declaration of 2-D slices
	dimensionsMap := make(map[int][]int) //Creating a map from int to slice of integers

	for i := 0; i < 2; i++ {
		color.Green("\nYou have to choose a proper dimension for MATRIX%d %s", i+1, " :-")
		j := 1
		for i := 0; i < 5; i++ {
			for k := 0; k < 5; k++ {
				a := []int{i + 1, k + 1}
				dimensionsMap[j] = a
				b := color.New(color.FgCyan)
				b.Print("[" + fmt.Sprint(j) + "]\t")
				m := color.New(color.FgMagenta)
				m.Print(fmt.Sprint(i+1) + "x" + fmt.Sprint(k+1) + "\t")
				j++
			}
			fmt.Println()
		}
		b := color.New(color.FgHiBlue)
		b.Print("\nPlease choose a proper dimension number for MATRIX", i+1, " => ")

		_, err := fmt.Scanf("%d", &option)

		if err != nil {
			Clear()
			color.Red("(Error) The entered dimension is not a valid integer.\n")
			i -= 1
			continue
		}
		if option > 25 || option < 1 {
			Clear()
			color.Red("(Error) The valid option number should be in range 1 ... 25.\n")
			i -= 1
			continue
		}

		if i == 0 { //Inputting Matrix1
			selectedMatrixCol = dimensionsMap[option][1]
			m1 = getMatrix(dimensionsMap[option], i+1)
		} else {
			if selectedMatrixCol != dimensionsMap[option][0] {
				Clear()
				color.Red("(Error) The number of columns in MATRIX1 & the number of rows in MATRIX2 should be same Or C1==R2\n")

				fmt.Println("Dimension( MATRIX1 ) : ", dimensionsMap[optionPrev][0], "x", dimensionsMap[optionPrev][1])
				i -= 1
				continue
			} else { //Inputting Matrix2
				fmt.Println("Dimension( MATRIX1 ) : ", dimensionsMap[optionPrev][0], "x", dimensionsMap[optionPrev][1])
				m2 = getMatrix(dimensionsMap[option], i+1)
			}
		}
		optionPrev = option
	}

	m3 := matrixMultiplication(m1, m2)
	Clear()
	color.Green("\nMatrix1(A) :-")
	showContents(m1)
	color.Green("\nMatrix2(B) :-")
	showContents(m2)
	color.Green("\nMatrix1 x Matrix2 (AxB):-")
	showContents(m3)
	color.Blue("\n\n----------------------------- End --------------------------------\n\n")
}

//A function that clears the screen(stdout)
func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//A function that returns a matrix after creating it according the dimesions parameter
func getMatrix(dimensions []int, matrixNumber int) [][]uint {
	color.Magenta("MATRIX" + fmt.Sprint(matrixNumber) + " (Rowwise entry) [" + fmt.Sprint(dimensions[0]) + "x" + fmt.Sprint(dimensions[1]) + "] :-")
	m := make([][]uint, dimensions[0])
	for i := 0; i < dimensions[0]; i++ {
		a := make([]uint, dimensions[1])
		for j := 0; j < dimensions[1]; j++ {
			b := color.New(color.Faint)
			b.Print("Enter Value( m1[" + fmt.Sprint(i) + "][" + fmt.Sprint(j) + "]) :  ")
			_, err := fmt.Scanf("%d", &a[j])
			if err != nil {
				var c string
				fmt.Scanf("%s", &c) //To feed new line in case of invalid entry...
				color.Red("(Error) This is not a valid integer value for integer Matrix.\n")
				j -= 1
				continue //It is not required here as it is the last statment...
			}
		}
		m[i] = a
	}
	return m
}

//A function that multiplies 2 matrices and returns the resulting one...
func matrixMultiplication(m1, m2 [][]uint) [][]uint {
	r1 := len(m1)
	c1 := len(m1[0])
	c2 := len(m2[0]) //r1==c1
	m3 := make([][]uint, r1)
	for i := 0; i < r1; i++ {
		m3[i] = make([]uint, c2)
	}

	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			m3[i][j] = 0
			for k := 0; k < c1; k++ {
				m3[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return m3
}

//A function that displays the matrix contents
func showContents(m [][]uint) {
	r := len(m)
	c := len(m[0])

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			magColor := color.New(color.FgHiCyan)
			magColor.Print(m[i][j], "\t")
		}
		fmt.Println()
	}
}
