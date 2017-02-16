/*
	{
		"date_of_creation" => "25 Dec 2016, Thurs",
		"aim_of_program"   => "To use structures with interfaces to organize informations",
		"coded_by"         => "Rishikesh Agrawani",
		"Go_version"	   => "1.7",
	}
*/

package main

import "fmt"

//Defining Teacher structure
type Teacher struct {
	name             string
	age              int
	teachingSubjects []string
}

//Defining Student structure
type Student struct {
	name        string
	age         int
	AllSubjects []string
}

//Defining a list of methods signature
type Details interface {
	getName() string       //returns a name of teacher/student
	getAge() int           //returns an age of person/teacher
	getSubjects() []string //returns a slice of languages related to teacher/student
}

//Teacher implements the getName()
func (t Teacher) getName() string {
	return t.name
}

//Teacher implements the getAge()
func (t Teacher) getAge() int {
	return t.age
}

//Teacher implements the getSubjects()
func (t Teacher) getSubjects() []string {
	return t.teachingSubjects
}

//Student implements the getName()
func (s Student) getName() string {
	return s.name
}

//Student implements the getAge()
func (s Student) getAge() int {
	return s.age
}

//Student implements the getSubjects()
func (s Student) getSubjects() []string {
	return s.AllSubjects
}

//Teacher & Student structs implement the Details interface, getDetails can accept both types of argument
func getDetails(person Details) {
	fmt.Println(person.getName())               //Calling a specific getName() according to type of person(Teacher/Student)
	fmt.Println(person.getAge())                //Calling a specific getAge() according to type of person(Teacher/Student)
	for i, book := range person.getSubjects() { //Calling a specific getSubjects() according to type of person(Teacher/Student)
		fmt.Println("Book", i+1, ":\t", book)
	}
	fmt.Println("\n")
}

func main() {
	teacher1 := Teacher{"Rob Pike", 59, []string{"Go", "C++"}}                              //Creating 1st teacher object
	teacher2 := Teacher{"Robert Griesemer", 52, []string{"Going Go", "The Journey Of C++"}} //Creating 2nd teacher object
	for _, teacher := range []Teacher{teacher1, teacher2} {                                 //Looping through list of teachers to extract the details of each
		getDetails(teacher)
	}

	student1 := Student{"Rishikesh Agrawani", 24, []string{"C", "C++", "C#", "Python", "Golang", "Core java", "PHP"}}       //Creating 1st student object
	student2 := Student{"Dennins Programmer", 35, []string{"Jumbo", "Dolly", "Limbo", "Ruby", "Javascript", "PHP", "Glue"}} //Creating 2nd student object
	student3 := Student{"Filodian Pellu", 26, []string{"Haskell", "Objecive C", "C++"}}                                     //Creating 3rd student object
	for _, student := range []Student{student1, student2, student3} {                                                       //Looping through list of students to extract the details of each
		getDetails(student)
	}
}

/*
Rob Pike
59
Book 1 :	 Go
Book 2 :	 C++


Robert Griesemer
52
Book 1 :	 Going Go
Book 2 :	 The Journey Of C++


Rishikesh Agrawani
24
Book 1 :	 C
Book 2 :	 C++
Book 3 :	 C#
Book 4 :	 Python
Book 5 :	 Golang
Book 6 :	 Core java
Book 7 :	 PHP


Dennins Programmer
35
Book 1 :	 Jumbo
Book 2 :	 Dolly
Book 3 :	 Limbo
Book 4 :	 Ruby
Book 5 :	 Javascript
Book 6 :	 PHP
Book 7 :	 Glue


Filodian Pellu
26
Book 1 :	 Haskell
Book 2 :	 Objecive C
Book 3 :	 C++

*/
