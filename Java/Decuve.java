import java.util.Scanner;

public class Decuve {

	static int poid;
	static int dose;
	static int heure = 0;
	static double decuve;
	static double alcool;
	static boolean homme;
	
	public static void main(String[] args) {
		
		@SuppressWarnings("resource")
		Scanner s = new Scanner(System.in);
		System.out.println("Etes vous un homme ? (true/false)");
		homme = s.nextBoolean();
		System.out.println("Quel est votre poid ? (kg)");
		poid = s.nextInt();
		System.out.println("Combien de verre avez vous bu ? (dose bar)");
		dose = s.nextInt();
		
		if(homme == true) {
			alcool = (dose*10)/(poid*0.7);
		}
		if(homme == false) {
			alcool = (dose*10)/(poid*0.6);
		}
		System.out.println("Vous avez " + alcool + "g/L d'alcool dans le sang");
		
		if(alcool >= 0.5) {
			System.out.println("Vous ne pouvez pas prendre le volant !");
			
			if(homme == true) {decuve = 0.085;}
			else {decuve = 0.1;}
			
			while(alcool >= 0.5) {alcool = alcool-decuve;heure++;}
			System.out.println("Vous pouvez reprendre le volant dans " + heure + "h");
		}
	}
}
