package main

import (
	"fmt"
)



func main() {

	var homme bool
	var poid int
	var dose int 
	var float int
	
	fmt.Println("Etes vous un homme ? (true/false)\n")
    fmt.Scan(&homme)
    fmt.Println("Quel est votre poid ? (kg)\n")
    fmt.Scan(&poid)
    fmt.Println("Combien de verre avez vous bu ? (dose bar)\n")
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
				continue
			} else {
				break
			}
		}

        fmt.Println("Vous pouvez reprendre le volant dans " + heure + "h")

	} else {

        fmt.Println("Vous pouvez prendre le volant.")
	}

}