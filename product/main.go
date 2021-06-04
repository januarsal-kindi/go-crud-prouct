package main

func main() {
	a := App{}
	a.Initialize(
		"root",
		"todo",
		"todo")

	a.Run(":8010")
}
