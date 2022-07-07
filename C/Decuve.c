#include <stdio.h>
#include <stdlib.h>

static int poid;
static int dose;
static int heure = 0;
static double decuve;
static double alcool;
static bool homme;

void main() {

    scanf("Etes vous un homme ? (true/false)", &homme);
    scanf("Quel est votre poid ? (kg)", &poid);
    scanf("Combien de verre avez vous bu ? (dose bar)", &dose);

    if(homme==true){
        alcool = (dose*10)/(poid*0.7);
    }
    if(homme==false) {
        alcool = (dose*10)/(poid*0.6);
    }
    printf("Vous avez", alcool, "g/L d'alcool dans le sang");

    if(alcool >= 0.5) {
        printf("Vous ne pouvez pas prendre le volant !")

        if(homme == true) {
            decuve = 0.085;
        }
        else {
            decuve = 0.1;
        }
        while(alcool >= 0.5) {
            alcool = alcool-decuve;
            heure++;
        }
        printf("Vous pouvez reprendre le volant dans ", heure, "h");

    }
}
