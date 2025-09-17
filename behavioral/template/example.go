package main

import "fmt"

// Beverage 饮料接口
type Beverage interface {
	BoilWater()
	Brew()
	PourInCup()
	CustomerWantsCondiments() bool
	AddCondiments()
	MakeBeverage()
}

// CaffeineBeverage 定义制作饮料的模板
type CaffeineBeverage struct {
	beverage Beverage
}

func (cb *CaffeineBeverage) MakeBeverage() {
	cb.beverage.BoilWater()
	cb.beverage.Brew()
	cb.beverage.PourInCup()

	if cb.beverage.CustomerWantsCondiments() {
		cb.beverage.AddCondiments()
	}
}

// Tea 具体的子类：茶
type Tea struct{}

func (t *Tea) BoilWater() {
	fmt.Println("Boiling water for tea.")
}

func (t *Tea) Brew() {
	fmt.Println("Steeping the tea.")
}

func (t *Tea) PourInCup() {
	fmt.Println("Pouring tea into cup.")
}

func (t *Tea) CustomerWantsCondiments() bool {
	var answer string
	fmt.Println("Do you want lemon with your tea (y/n)?")
	_, err := fmt.Scanln(&answer)
	if err != nil {
		return false
	}
	return answer == "y"
}

func (t *Tea) AddCondiments() {
	fmt.Println("Adding lemon to tea.")
}

func (t *Tea) MakeBeverage() {
	cb := &CaffeineBeverage{beverage: t}
	cb.MakeBeverage()
}

// Coffee 具体的子类：咖啡
type Coffee struct{}

func (c *Coffee) BoilWater() {
	fmt.Println("Boiling water for coffee.")
}

func (c *Coffee) Brew() {
	fmt.Println("Brewing the coffee.")
}

func (c *Coffee) PourInCup() {
	fmt.Println("Pouring coffee into cup.")
}

func (c *Coffee) CustomerWantsCondiments() bool {
	var answer string
	fmt.Println("Do you want milk and sugar with your coffee (y/n)?")
	_, err := fmt.Scanln(&answer)
	if err != nil {
		return false
	}
	return answer == "y"
}

func (c *Coffee) AddCondiments() {
	fmt.Println("Adding milk and sugar to coffee.")
}

func (c *Coffee) MakeBeverage() {
	cb := &CaffeineBeverage{beverage: c}
	cb.MakeBeverage()
}

func main() {
	// 测试：制作茶
	tea := &Tea{}
	tea.MakeBeverage()

	// 测试：制作咖啡
	coffee := &Coffee{}
	coffee.MakeBeverage()
}
