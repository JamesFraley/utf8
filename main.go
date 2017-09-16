package main

import "fmt"
import "./printStuff"

func main() {
	var myString = `
	<!DOCKTYPE html>
	<html lang="en">
	<body>
	<meta charset="UTF-8"
	Hello
	</body>`

   fmt.Println("hello")
   printStuff.MyPrint(myString)
}
