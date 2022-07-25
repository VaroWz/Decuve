#include <stdio.h>

static int poid;
static int dose;
static int heure = 0;
static double decuve;
static double alcool;
static char homme;

int main() {

    //printf("Lancement du programme... En C");

    //Sans boolean je peux vérifier par un char
    printf("Etes vous un homme ? (true/false)\n");
    scanf("%c", &homme);
    printf("Quel est votre poid ? (kg)\n");
    scanf("%i", &poid);
    printf("Combien de verre avez vous bu ? (dose bar)\n");
    scanf("%i", &dose);

    if(homme==1) {
        alcool = (dose*10)/(poid*0.7);
    }
    if(homme==1) {
        alcool = (dose*10)/(poid*0.6);
    }
    printf("Vous avez", alcool, "g/L d'alcool dans le sang");

    if(alcool >= 0.5) {
        printf("Vous ne pouvez pas prendre le volant !");

        if(homme == 1) {
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
