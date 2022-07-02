package main

import (
	"fmt"
)



func main() {

	var homme bool = false
	var poid float64
	var dose float64
	var alcool float64
	var heure float64 = 0
	var decuve float64
	
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