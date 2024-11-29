package main

import (
	"fmt"
	"errors"
)

func dogFoodAmount(amount int) float64 {
    return 10.0 * float64(amount)
}

func catFoodAmount(amount int) float64 {
    return 5.0 * float64(amount)
}

func hamsterFoodAmount(amount int) float64 {
    return 0.250 * float64(amount)
}

func spiderFoodAmount(amount int) float64 {
    return 0.150 * float64(amount)
}

func animal(animalType string) (func(int) float64, error) {
    animals := map[string]func(int) float64{"dog": dogFoodAmount, "cat": catFoodAmount, "hamster": hamsterFoodAmount, "spider": spiderFoodAmount}
    if _, ok := animals[animalType]; !ok {
        return nil, errors.New("Not a valid animal")
    }
    return animals[animalType], nil
}

func main() {
    animalDog, msg := animal("dog")
    if msg != nil {
        return
    }
    animalCat, msg := animal("cat")
    if msg != nil {
        return
    }
    animalHamster, msg := animal("hamster")
    if msg != nil {
        return
    }
    animalSpider, msg := animal("spider")
    if msg != nil {
        return
    }

    var sum float64
    amount := animalDog(10)
    sum += amount
    fmt.Println("Dog - Amount: 10 - Qtd de Alimento", amount, "kg")
    amount = animalCat(18)
    sum += amount
    fmt.Println("Cat - Amount: 18 - Qtd de Alimento", amount, "kg")
    amount = animalHamster(20)
    sum += amount
    fmt.Println("Hamster - Amount: 20 - Qtd de Alimento", amount, "kg")
    amount = animalSpider(14)
    sum += amount
    fmt.Println("Spider - Amount: 14 - Qtd de Alimento", amount, "kg")

    fmt.Println("Total Amount for 14 Spiders + 20 Hamsters + 18 Cats + 10 Dogs", sum, "kg")
}