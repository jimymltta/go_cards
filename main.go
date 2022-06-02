// var card string = "String"
// card := "String"

// ces deux initialisations de variables sont la même chose, juste un raccourci
// une fois initialiséé avec :=, on peut assigner une nouvelle valeur juste avec =

// les arrays et les slices sont tous deux des tableaux de valeurs, mais il y a une différence
// un array a un nombre d'éléments fixe, alors qu'un slice est "extensible" (comme un array en JS)

// pour compiler le code Go, commande go run main.go. Crée un nouveau ficher go.exe ou exec qui est un éxécutable

package main // obligatoire, spécifie que c'est le fichier principal à compiler

import "fmt" // import de package(s)

func main() { // Cette fonction est obligatoire, c'est la fonction principale du programme
	cards := []string{"Ace of Diamonds", newCard()} // c'est un slice = comme un array mais auquel on peut ajouter/enlever des éléments (voir haut de page)
	cards = append(cards, "Six of Spades") // ne modifie pas le slice, mais en crée un nouveau et l'assigne à cards

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}