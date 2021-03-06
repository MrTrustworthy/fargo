package entities

import (
	"errors"
	"strconv"
)

type Item struct {
	Name string
}

type Inventory map[Item]int

func (i *Inventory) IsEmpty() bool {
	return len(*i) == 0
}

func (i Inventory) Clear() {
	for k := range i {
		delete(i, k)
	}
}

func (i Inventory) Add(item Item, count int) {
	i[item] += count
}

func (i Inventory) RemoveAll(item Item) (int, error) {
	amount := i[item]
	return amount, i.Remove(item, i[item])
}

func (i Inventory) Remove(item Item, count int) error {
	if i[item] > count {
		return errors.New("not enough items to remove")
	}
	i[item] -= count
	if i[item] == 0 {
		delete(i, item)
	}
	return nil
}

func (i Inventory) String() string {
	s := "Items: "
	for key, val := range i {
		s += " [" + strconv.Itoa(val) + "] " + key.Name
	}
	return s
}

func (i Inventory) Size() int {
	sum := 0
	for _, value := range i {
		sum += value
	}
	return sum
}

func (i Inventory) FillFrom(other Inventory) {
	for key, value := range other {
		i.Add(key, value)
	}
	other.Clear()
}

func NewSampleInventory() *Inventory {
	item1, item2 := Item{Name: "Gold"}, Item{Name: "Silver"}
	inventory := &Inventory{}
	inventory.Add(item1, 1)
	inventory.Add(item2, 3)
	return inventory
}
