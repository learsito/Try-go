package main

import "fmt"

func main() {
	var uno int
	var dos int
	var tres int

	fmt.Println("Ingresa un número: ")
	fmt.Scanln(&uno)

	fmt.Println("Ingresa otro número: ")
	fmt.Scanln(&dos)

	fmt.Println("Ingresa el último número: ")
	fmt.Scanln(&tres)

	fmt.Println("El número mayor es: ")

	if (uno > dos) && (uno > tres) {
		fmt.Println(uno)
	} else if (dos > uno) && (dos > tres) {
		fmt.Println(dos)
	} else if (tres > uno) && (tres > dos) {
		fmt.Println(tres)
	} else {
		fmt.Println("Los números son iguales")
	}

}
