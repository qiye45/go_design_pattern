package main

import "fmt"

// Button 两类产品接口
type Button interface{ Click() }
type TextBox interface{ Show() }

// UIFactory 抽象工厂
type UIFactory interface {
	CreateButton() Button
	CreateTextBox() TextBox
}

// WinButton win系列
type WinButton struct{}

func (WinButton) Click() { fmt.Println("Win按钮") }

type WinTextBox struct{}

func (WinTextBox) Show() { fmt.Println("Win文本框") }

type WinFactory struct{}

func (WinFactory) CreateButton() Button   { return WinButton{} }
func (WinFactory) CreateTextBox() TextBox { return WinTextBox{} }

func main() {
	var f UIFactory = WinFactory{}
	f.CreateButton().Click()
	f.CreateTextBox().Show()
}
