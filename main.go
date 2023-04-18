package main

import (
	"fmt"
)

func main() {
	fmt.Println("Приветствую в игре.")

	level := 1
	expirience := 0

	fmt.Println("Ваш уровень:", level)
	fmt.Println("Ваш опыт:", expirience)

	fmt.Println("Сколько очков опыта вы получили на последнем задании?")
	fmt.Scan(&expirience)

	if expirience >= 5000 {
		level = 4
		fmt.Println("Ваш уровень:", level)
		fmt.Println("Ваш опыт:", expirience)
	} else if expirience >= 2500 {
		level = 3
		fmt.Println("Ваш уровень:", level)
		fmt.Println("Ваш опыт:", expirience)
	} else if expirience >= 1000 {
		level = 2
		fmt.Println("Ваш уровень:", level)
		fmt.Println("Ваш опыт:", expirience)
	} else {
		level = 1
		fmt.Println("Ваш уровень:", level)
		fmt.Println("Ваш опыт:", expirience)
		fmt.Println("Чтобы перейти на уровень 2, вам нужно достичь отметки в 1000 очков.")
	}
}
