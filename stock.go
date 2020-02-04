package stockpile

import "fmt"

type Stockpile []Stock

type Stock struct {
	Name     string
	Quantity int
}

func NewStockpile() Stockpile {
	return Stockpile{}
}

func (s Stockpile) String() string {
	var res string
	for _, v := range s {
		res += fmt.Sprintf("%-30s %5d\n", v.Name, v.Quantity)
	}
	return res
}

func (s *Stockpile) Add(stock Stock) error {
	if stock.Name == "" {
		return fmt.Errorf("trying to add a stock with no name")
	}
	if stock.Quantity <= 0 {
		return fmt.Errorf("trying to add a stock with 0 or less quantity")
	}
	added := false
	for k, v := range *s {
		if v.Name == stock.Name {
			(*s)[k].Quantity += stock.Quantity
			added = true
		}
	}
	if !added {
		*s = append(*s, stock)
	}
	return nil
}
