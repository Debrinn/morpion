package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var (
	coups int = 0
	actualPlayer int = 1
	morpion = [9]string{
		"1","2","3",
		"4","5","6",
		"7","8","9"}
	end bool=true
)
func main() {
	afficherMorpion(morpion)
	for end{
		scanner :=bufio.NewScanner(os.Stdin)
		fmt.Print("Joueur numéro ",actualPlayer, " vous pouvez jouer : ")//numéro à changer
		scanner.Scan()
		if scanner.Text()=="stop"{
			break
		}
		value, err := strconv.Atoi(scanner.Text())
		if err !=nil{
			fmt.Println("It's NOT a number")
			continue
		}
		if value>9 || value<1{
			fmt.Println("Valeur impossible")
			continue
		}else if morpion[value-1]=="X" || morpion[value-1]=="O"{
			fmt.Println("Un joueur a déja joué ici")
			continue
		}else{
			if actualPlayer==1{
				morpion[value-1]="X"
			}else{
				morpion[value-1]="O"
			}

			afficherMorpion(morpion)
			actualPlayer=2
			}

		coups++
		var compteur int=0
		for i:=1;i<8;i+=3{

			var tableau1 =[3]string{
				morpion[i-1],
				morpion[i],
				morpion[i+1],
			}

			if tableau1[0]==tableau1[1] && tableau1[1]==tableau1[2]{
				if tableau1[0]=="X"{
					fmt.Println("Le joueur 1 a gagné")
				}else{
					fmt.Println("Le joueur 2 a gagné")
				}
				end=false
			}
			if morpion[compteur]==morpion[compteur+3] && morpion[compteur+3]==morpion[compteur+6]{
				if morpion[compteur]=="X"{
					fmt.Println("Le joueur 1 a gagné")
				}else{
					fmt.Println("Le joueur 2 a gagné")
				}
				end=false
			}
			compteur++
		}
		if (morpion[0]==morpion[8] && morpion[8]==morpion[4]) || (morpion[2]==morpion[4] && morpion[4]==morpion[6]){
			if morpion[4]=="X"{
				fmt.Println("Le joueur 1 a gagné")
			}else{
				fmt.Println("Le joueur 2 a gagné")
			}
			end=false
		}
		if coups==9{
			fmt.Println("Egalité")
			end=false
		}
	}
}
func afficherMorpion(tableau[9]string){
	for i,val := range tableau{
		if (i+1)%3==0{
			fmt.Println(val)
		} else{
			fmt.Print(val," ")
		}
	}
}
