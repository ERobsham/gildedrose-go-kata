package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/erobsham/gildedrose-go-kata/gildedrose"
)

func main() {
	fmt.Println("OMGHAI!")

	var items = []*gildedrose.Item{
		{Name: "+5 Dexterity Vest", SellIn: 10, Quality: 20},
		{Name: "Aged Brie", SellIn: 2, Quality: 0},
		{Name: "Elixir of the Mongoose", SellIn: 5, Quality: 7},
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: 0, Quality: 80},
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: -1, Quality: 80},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 49},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 49},
		{Name: "Conjured Mana Cake", SellIn: 3, Quality: 6}, // <-- :O
	}

	days := 2
	var err error
	if len(os.Args) > 1 {
		days, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		days++
	}

	fmt.Printf("-------- initial state --------\n")
	logItems(items)

	for day := 0; day < days; day++ {
		gildedrose.UpdateQuality(items)

		fmt.Printf("-------- post day %d --------\n", day)
		logItems(items)
	}
}

func logItems(items []*gildedrose.Item) {
	fmt.Println("| SellIn | Quality | Name")
	for i := 0; i < len(items); i++ {
		fmt.Printf("| % 4d   | % 4d    | %v \n", items[i].SellIn, items[i].Quality, items[i].Name)
	}
	fmt.Println("")
}
