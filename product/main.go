package main

func main() {
	a := App{}
	a.Initialize(
		"localhost",
		5432,
		"postgres",
		"todo",
		"todo")

	a.Run(":8010")
}
