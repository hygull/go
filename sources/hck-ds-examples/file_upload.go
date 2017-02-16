package main

import "github.com/kataras/iris"
import "os"
import "io"
import "time"
import "fmt"

func UploadFile(ctx *iris.Context) {
	ctx.SetContentType("application/json")
	// Get name and email
	// name := ctx.FormValueString("name")
	// email := ctx.FormValueString("email")
	// Get avatar
	fmt.Print("\nGetting file...")
	avatar, err := ctx.FormFile("file")
	if err != nil {
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}
	fmt.Print("\nSuccessfully got the file...")
	// Source

	fmt.Print("\nCreating new source file...")
	src, err := avatar.Open()
	if err != nil {
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}
	defer src.Close()

	// Destination
	fmt.Print("\nCreating destination file")
	avatar.Filename = time.Now().String()[0:19] + avatar.Filename
	FileUploadPathFromAppRoot := "./upload/"
	MkDir(FileUploadPathFromAppRoot)

	fmt.Print("File name : ", avatar.Filename, "\n")
	dst, err := os.Create(FileUploadPathFromAppRoot + avatar.Filename)
	//dst, err := os.Create(avatar.Filename)
	//dst, err := os.Create(up_river_hill.jpeg)
	fmt.Printf("\n%T...%T...%T", dst, avatar, avatar.Filename)
	if err != nil {
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}
	defer dst.Close()
	fmt.Print("\nGoing to copy source file to destination")
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}
	fmt.Print("File successfully created...")
	fmt.Print("\nDisplaying JSON Response\n")
	ctx.JSON(200, iris.Map{"status": 200, "msg": "file successfully uploaded"})
	//ctx.HTML(iris.StatusOK, "<b>Thanks!</b>")
}

func MkDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0777)
		fmt.Println("Directory " + dir + " created.")
	} else {
		fmt.Println("Directory " + dir + " exists.")
	}
}
