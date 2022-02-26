package main

import "fmt"

func main() {
	var p Price = 2595
	fmt.Println(p.String())

	userCart := Cart{Items: []string{}, TotalPrice: p}
	userCart.AddItem("eggs")
	userCart.AddItem("milk")
	userCart.AddItem("chocolate")
	fmt.Println(userCart)
}

// Price is the cost of something in US cents.
type Price int64

// String is the string representation of a Price
// These should be represented in US Dollars
// Example: 2595 cents => $25.95
func (p Price) String() string {
	return fmt.Sprintf("$%f", p/100.0)
}

// Prices is a map from an item to its price.
var Prices = map[string]Price{
	"eggs":          219,
	"bread":         199,
	"milk":          295,
	"peanut butter": 445,
	"chocolate":     150,
}

// RegisterItem adds the new item in the prices map.
// If the item is already in the prices map, a warning should be displayed to the user,
// but the value should be overwritten.
// Bonus (1pt) - Use the "log" package to print the error to the user
func RegisterItem(prices map[string]Price, item string, price Price) {
	boo := false
	for k, _ := range prices {
		if k == item {
			boo = true
			break
		}
	}
	if !boo {
		prices[item] = price
	} else {
		fmt.Print("Error: ", item, " is occur in prices map")
		prices[item] = price
	}
}

// Cart is a struct representing a shopping cart of items.
type Cart struct {
	Items      []string
	TotalPrice Price
}

// hasMilk returns whether the shopping cart has "milk".
func (c *Cart) hasMilk() bool {
	return c.HasItem("milk")
}

// HasItem returns whether the shopping cart has the provided item name.
func (c *Cart) HasItem(item string) bool {
	for i := 0; i < len(c.Items); i++ {
		if item == c.Items[i] {
			return true
		}
	}
	return false
}

// AddItem adds the provided item to the cart and update the cart balance.
// If item is not found in the prices map, then do not add it and print an error.
// Bonus (1pt) - Use the "log" package to print the error to the user
func (c *Cart) AddItem(item string) {
	boo := false
	for k, v := range Prices {
		if k == item {
			c.Items = append(c.Items, item)
			c.TotalPrice -= v
			boo = true
		}
	}
	if !boo {
		fmt.Print("Error", item, " not occur in Prices map")
	}
}

// Checkout displays the final cart balance and clears the cart completely.
func (c *Cart) Checkout() {
	fmt.Print("Cart balance:", c.TotalPrice)
	c.Items = nil
	c.TotalPrice = 0
}
