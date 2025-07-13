package main

import "fmt"

func main() {
	var size int

	fmt.Print("Введите размер шахматной доски: ")
	_, err := fmt.Scan(&size)

	if err != nil {
		fmt.Println("Ошибка", err)
		return
	}
	//fmt.Println(sizeOfDesk)
	for i := 0; i < size; i += 1 {

		for j := 0; j < size; j += 1 {

			if (i+j)%2 == 0 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()

	}

}
