
/*
GO


Créé en 2009
Influencé par C / Unix

Pourquoi GO ?**
-> Performances

**Que faire avec GO ?**
-> API
-> Tooling

**Go** :
- Statique
- Compilé
- C-like
- GC
- Multi-cpu/parallelisme
- Un fichier binaire après compilation
- Multiplateforme

**Clients** :
Dropbox, github, Netflix, molotovtv, leboncoin


*/

import "fmt"
import("fmt"
	 	"string")

func blabla() {}

/*
bool
string
int/int8- 16-32-64
uint/uint8- 16-32-64
byte
float32
float64
Rune
complex64/128 (nombres complexes)

-> orienté math et performance
-> inférence de types
-> statique
*/

var age int = 20
var age = 20
var age int // 0
var name String = "bob"
var name String // empty string «  »
var weight, height int = 20
var weight, height int = 200, 289

test := "pouet"

//Possibilité de faire des variables globales :

var pi = 3.14
func main() {}

//Visibilité :

var private
var Public

if age > 2 {}


switch age {
	case 1 : blabla
}

switch {
	case age == 2 : blabla
}

/*
Opérations impossibles :
int + string // erreur de compil (sauf en javascript lol)
idem int + float // :/

Conversion de type :
**upcasté** la valeur entiere (int) pour etre en float car on ne perd pas d’info d’un int a un float, l’inverse non.
int = int + float // non
float = float + int // oui
*/

var percent float64 = 2.12
var toInt int = int(percent) // 2

percent = percent + 34
//ça marche car upcasted pour etre en float

n := 42
f := float(n) + 0.42


//Visibilité comme pour les variables

func FunctionPublic() {}
func functionPrivee() {}

func nomdeFonction(name string) {}
func retourFunction(name string) bool {}
func somme(x, y, z int) (sum int) {
//la var de retour sum sera déjà declarée
	sum x + y+ z
	return sum
}

//retour de fonction multiple
func MyFunc()(string, int) {
	return "jean", 9
}

name, age := MyFunc()


/*
Tableaux
Index à zero, taille définitive
initialisé a la valeur vide du type

Slice : tranche d’un tableau
Taille dynamique
*/

var names [3]string
fmt.Printf(« Names=%v (len=%v), names, len(names) »

//slice
s := make([]type, taille, capacité du tableau derriere)
s := make([]int, 3)
s = append(s, 12)
//creatin d’un nouveau tableau plus grand avec la meme valeur que l’ancien, invisible, comme si on avait juste augmenté la taille du tableau

s := make([]int, 2, 3)
nums[0] = 10
nums[1] = -2
nums[2] = 5 // on ajoute une valeur au slice au dela de sa taille initiale, le tableau grandit
//si on ne precise pas la taille du tableau elle est implicite
letters := []string{« a », « a », « a »}
//longueur et capacité de 3 et longueur 3
letters = append(letters, « e »)
//la longueur passe a 4 pour le slice et 6 pour le tableau

//subslice
sub1 := lettesr[0:2]
sub1 := lettesr[0:]//jusqu’au dernier index
sub1 := lettesr[:3]//du premier au 3e indice

//declaration
var a [2]int // tableau
var a []int // slice

//Les boucles

event := 0
for event > 0 {
	//blabla
	event ++
	continue // ou break
 	}

names := []string{«a », «a », «a »}
for i, n := range names {
		//blabla
 	}

for i, c := range «bla» {
		//blabla
 	}

//Erreurs

/*
-> early return pour plus de visibilité
tester en premier le cas d’erreur avant le reste
*/

func myfunc()(int, error) {}

//renvoyer une erreur
errors.New("Error")
