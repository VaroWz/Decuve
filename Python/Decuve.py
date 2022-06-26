heure = 0

def main():

    #Demande du genre
    homme = bool(input("Etes vous un homme ? (true/false)\n"))
    #Demande du poid
    poid = int(input("Quel est votre poid ? (kg)\n"))
    #Demande de dose
    dose = int(input("Combien de verre avez vous bu ? (dose bar)\n"))

    if homme == True:

        #taux d'alcool sans décuve
        alcool = (dose * 10) / (poid * 0.7);

    if homme == False:

        #taux d'alcool sans décuve
        alcool = (dose * 10) / (poid * 0.6);

    print("Vous avez " + str(alcool) + "g/L d'alcool dans le sang")

    if alcool >= 0.5:

        print("Vous ne pouvez pas prendre le volant !")

        if homme == True:
            decuve = 0.085

        else:
            decuve = 0.1

        while alcool >= 0.5:

            alcool = alcool-decuve
            heure = heure+1

        print("Vous pouvez reprendre le volant dans " + str(heure) + "h")

    else:

        print("Vous pouvez prendre le volant.")

if __name__ == '__main__':

    main()