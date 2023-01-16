package main

func main() {
	slice := []int{1, 2, 3}
	for i, v := range slice {
		println(i, v)
	}
	port := map[string]int{"https": 18, "http": 19}
	for k, val := range port {
		println(k, val)
	}
	ports := map[string]int{"ash": 1, "htt": 119}
	for k, _ := range ports {
		println(k)
	}
}
