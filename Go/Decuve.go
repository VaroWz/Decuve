package main

import (
	"fmt"
)
var(
	homme bool = false
	poid float64
	dose float64
	alcool float64
	heure float64 = 0
	decuve float64
)
func main() {
	
	fmt.Println("Etes vous un homme ? (true/false)")
    fmt.Scan(&homme)
    fmt.Println("Quel est votre poid ? (kg)")
    fmt.Scan(&poid)
    fmt.Println("Combien de verre avez vous bu ? (dose bar)")
	fmt.Scan(&dose)

    if homme == true {

        alcool = (dose * 10) / (poid * 0.7)
	}
    if homme == false {

        alcool = (dose * 10) / (poid * 0.6)
	}
    fmt.Println("Vous avez", alcool, "g/L d'alcool dans le sang")

    if alcool >= 0.5 {

        fmt.Println("Vous ne pouvez pas prendre le volant !")

        if homme == true {
            decuve = 0.085
		} else {
            decuve = 0.1
		}
		for true {
			if alcool >= 0.5 {
				alcool = alcool-decuve
				heure++
				continue
			} else {
				break
			}
		}

        fmt.Println("Vous pouvez reprendre le volant dans", heure, "heures")

	} else {

        fmt.Println("Vous pouvez prendre le volant.")
	}

}