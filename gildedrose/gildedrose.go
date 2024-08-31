package gildedrose

import (
	"cmp"
	"strings"
)

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		item.updateQuality()
		item.updateSellIn()
	}
}

type ItemType int

const (
	Type_Normal ItemType = iota
	Type_Aged
	Type_Legendary
	Type_BackstagePasses
)

func (item *Item) intoType() ItemType {
	var legendaryItems = map[string]struct{}{
		"Sulfuras, Hand of Ragnaros": {},
	}

	var agedItems = map[string]struct{}{
		// aged cheeses / fine wines / etc
		"Aged Brie": {},
	}

	if _, exists := legendaryItems[item.Name]; exists {
		return Type_Legendary
	} else if _, exists := agedItems[item.Name]; exists {
		return Type_Aged
	} else if strings.HasPrefix(item.Name, "Backstage passes") {
		return Type_BackstagePasses
	} else {
		return Type_Normal
	}
}

func (item *Item) updateQuality() {
	switch item.intoType() {
	case Type_Legendary:
		updateQuality_Legendary(item)
	case Type_Aged:
		updateQuality_Aged(item)
	case Type_BackstagePasses:
		updateQuality_BackstagePass(item)
	case Type_Normal:
		updateQuality_Normal(item)
	}
}

func (item *Item) updateSellIn() {
	if item.intoType() == Type_Legendary {
		return
	} else {
		item.SellIn -= 1
	}
}

func (i *Item) adjustQualityBy(delta int) {
	i.Quality = clamp(i.Quality+delta, 0, 50)
}

func updateQuality_Legendary(item *Item) {}
func updateQuality_Normal(item *Item) {
	if item.SellIn > 0 {
		item.adjustQualityBy(-1)
	} else {
		item.adjustQualityBy(-2)
	}
}
func updateQuality_Aged(item *Item) {
	if item.SellIn > 0 {
		item.adjustQualityBy(1)
	} else {
		item.adjustQualityBy(2)
	}
}
func updateQuality_BackstagePass(item *Item) {
	if item.SellIn < 1 {
		item.Quality = 0
	} else if item.SellIn < 6 {
		item.adjustQualityBy(3)
	} else if item.SellIn < 11 {
		item.adjustQualityBy(2)
	} else {
		item.adjustQualityBy(1)
	}
}

// region -- misc helpers

func clamp[T cmp.Ordered](n T, low T, high T) T {
	return max(min(n, high), low)
}

//endregion -- misc helpers
