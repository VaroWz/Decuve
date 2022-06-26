def main():

    homme = bool(input("Etes vous un homme ? (true/false) "))
    poid = int(input("Quel est votre poid ? (kg) "))
    dose = int(input("Combien de verre avez vous bu ? (dose bar) "))

    if homme == True:
        alcool = (dose * 10) / (poid * 0.7);

    if homme == False:
        alcool = (dose * 10) / (poid * 0.6);

    print("Vous avez " + str(alcool) + "g/L d'alcool dans le sang")

    if alcool >= 0.5:
        print("Vous ne pouvez pas prendre le volant !")

        if homme == True:
            decuve = 0.085

        if homme == False:
            decuve = 0.1

        heure = 0
        while alcool >= 0.5:

            alcool = alcool-decuve
            heure = heure+1

        print("Vous pouvez reprendre le volant dans " + str(heure) + "h")

if __name__ == '__main__':
    main()