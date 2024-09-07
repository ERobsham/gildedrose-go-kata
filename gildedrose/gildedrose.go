package gildedrose

import (
	"cmp"
	"strings"
)

//region -- Constants / Definitions

const (
	Quality_Min = 0
	Quality_Max = 50

	Prefix_BackstagePass = "Backstage passes"
	Prefix_Conjured      = "Conjured"
)

var legendaryItems = map[string]struct{}{
	"Sulfuras, Hand of Ragnaros": {},
}

var agedItems = map[string]struct{}{
	// aged cheeses / fine wines / etc
	"Aged Brie": {},
}

//endregion -- Constants / Definitions

//region -- Base Functionality

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		u := UpdaterFor(item)
		u.UpdateQuality(item)
		u.UpdateSellIn(item)
	}
}

//endregion -- Base Functionality

//region -- ItemType specific behavior

func UpdaterFor(item *Item) ItemUpdater {
	var normal NormalItem
	updater := ItemUpdater{
		UpdateQuality: normal.UpdateQuality,
		UpdateSellIn:  normal.UpdateSellIn,
	}

	if _, exists := legendaryItems[item.Name]; exists {
		var legendary LegendaryItem
		updater.UpdateQuality = legendary.UpdateQuality
		updater.UpdateSellIn = legendary.UpdateSellIn
	} else if _, exists := agedItems[item.Name]; exists {
		var aged AgedItem
		updater.UpdateQuality = aged.UpdateQuality
	} else if strings.HasPrefix(item.Name, Prefix_BackstagePass) {
		var backstage BackstagePassItem
		updater.UpdateQuality = backstage.UpdateQuality
	} else if strings.HasPrefix(item.Name, Prefix_Conjured) {
		var conjured ConjuredItem
		updater.UpdateQuality = conjured.UpdateQuality
	}

	return updater
}

type ItemUpdater struct {
	UpdateQuality func(item *Item)
	UpdateSellIn  func(item *Item)
}

type LegendaryItem struct{}
type NormalItem struct{}
type AgedItem struct{}
type BackstagePassItem struct{}
type ConjuredItem struct{}

func (NormalItem) UpdateSellIn(item *Item)    { item.SellIn -= 1 }
func (LegendaryItem) UpdateSellIn(item *Item) {}

func (NormalItem) UpdateQuality(item *Item) {
	if item.SellIn > 0 {
		item.adjustQualityBy(-1)
	} else {
		item.adjustQualityBy(-2)
	}
}
func (AgedItem) UpdateQuality(item *Item) {
	if item.SellIn > 0 {
		item.adjustQualityBy(1)
	} else {
		item.adjustQualityBy(2)
	}
}
func (BackstagePassItem) UpdateQuality(item *Item) {
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
func (ConjuredItem) UpdateQuality(item *Item) {
	normal := NormalItem{}
	normal.UpdateQuality(item)
	normal.UpdateQuality(item)
}
func (LegendaryItem) UpdateQuality(item *Item) {}

//endregion -- ItemType specific behavior

//region -- misc helpers

func (i *Item) adjustQualityBy(delta int) {
	i.Quality = clamp(i.Quality+delta, Quality_Min, Quality_Max)
}

func clamp[T cmp.Ordered](n T, low T, high T) T {
	return max(min(n, high), low)
}

//endregion -- misc helpers
