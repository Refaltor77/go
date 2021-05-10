package main

import "fmt"

func main()  {

	//je déclare mes variables avec le type
	var entier int = 10 // int
	var float float32 = 10.50 // float
	var text string = "salut, je suis un texte" // string
	var vrai bool = true // boolean

	// typage dynamique
	floatSansLeDire := 1.05 // pas besoin de dire que c'est un float !


	// constante (valeur fixe)
	const uneConst int = 1


	// J'affiche les variables
	fmt.Println(entier)
	fmt.Println(float)
	fmt.Println(text)
	fmt.Println(vrai)
	fmt.Println(floatSansLeDire)
	fmt.Println(uneConst)
	fmt.Println("------------------")


	// les calcules
	a := 10
	b := 1
	fmt.Println("resulat du calcul addition: ", a + b)
	fmt.Println("resulat du calcul multiplication: ", a * b)
	fmt.Println("resulat du calcul division: ", a / b)
	fmt.Println("resulat du calcul soustraction: ", a - b)

	// incrémentation
	c := 1
	c++

	// décrémentation
	c--

	// changer le type d'une variable temporairement
	var changement int = int(float) // ça va supprimer la virgule après le 10 ducoup
	fmt.Println(changement)

}