package main

type Student struct {
	Name string
}

func main() {
	s := Student{Name: "Manzil"}
	switch s.Name {
	case "Manzil":
		println("update manzil")
	case "ZerOng":
		println("update ZerOng")
	case "Aspak":
		println("Get Aspak")
	case "Prasant":
		println("Update Prasant")
	default:
		println("Unknown Name")
	}
}
