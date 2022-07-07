#include <stdio.h>
#include <stdlib.h>

static int poid;
static int dose;
static int heure = 0;
static double decuve;
static double alcool;
static boolean homme;

void main() {

    scanf("Etes vous un homme ? (true/false)", &homme);
    scanf("Quel est votre poid ? (kg)", &poid);
    scanf("Combien de verre avez vous bu ? (dose bar)", &dose);

    if(homme==true){
        alcool = (dose*10)/(poid*0.7);
        printf(alcool);
    }


}
