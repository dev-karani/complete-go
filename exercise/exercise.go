package main

import "fmt"

func main() {

}

type Player struct {
	Name      string
	Inventory []Item
}
type Item struct {
	Name string
	Type string
}

func (p *Player) PickUpItem(item Item) {
	p.Inventory := append(p.Inventory, item)
	fmt.Printf("%s picked an item %v\n", p.Name, item.Name)
}
func (p *Player) DropItem(item Item) {
	deletedItem := delete(p.Inventory, item.Name)
}
func (p *Player) UseItem(item Item) {
	fmt.Println("something", item)
}