package gildedrose

import (
	"reflect"
	"testing"
)

//region -- Regression Tests

func Test_Standard_UpdateQuality_Rules(t *testing.T) {
	tests := []struct {
		name     string
		args     []*Item
		expected []*Item
	}{
		{
			name: "Quality and SellIn reduce by one each day",
			args: []*Item{
				{Name: "Rabbit's Foot", Quality: 10, SellIn: 10},
				{Name: "Vibrant Plume", Quality: 1, SellIn: 1},
			},
			expected: []*Item{
				{Name: "Rabbit's Foot", Quality: 9, SellIn: 9},
				{Name: "Vibrant Plume", Quality: 0, SellIn: 0},
			},
		},
		{
			name: "When SellIn is less than 1 reduce Quality twice as fast",
			args: []*Item{
				{Name: "Rabbit's Foot", Quality: 10, SellIn: 0},
				{Name: "Vibrant Plume", Quality: 2, SellIn: 0},
			},
			expected: []*Item{
				{Name: "Rabbit's Foot", Quality: 8, SellIn: -1},
				{Name: "Vibrant Plume", Quality: 0, SellIn: -1},
			},
		},
		{
			name: "Quality can never go negative",
			args: []*Item{
				{Name: "Rabbit's Foot", Quality: 0, SellIn: 10},
				{Name: "Vibrant Plume", Quality: 0, SellIn: 0},
			},
			expected: []*Item{
				{Name: "Rabbit's Foot", Quality: 0, SellIn: 9},
				{Name: "Vibrant Plume", Quality: 0, SellIn: -1},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateQuality(tt.args)

			for idx, got := range tt.args {
				want := tt.expected[idx]

				if !reflect.DeepEqual(got, want) {
					t.Errorf("got: %+v  want:%+v", got, want)
				}
			}
		})
	}
}
func Test_AgedBrie_UpdateQuality_Rules(t *testing.T) {
	tests := []struct {
		name     string
		args     []*Item
		expected []*Item
	}{
		{
			name: "Aged Brie increases in Quality each day",
			args: []*Item{
				{Name: "Aged Brie", Quality: 10, SellIn: 10},
				{Name: "Aged Brie", Quality: 10, SellIn: 0},
			},
			expected: []*Item{
				{Name: "Aged Brie", Quality: 11, SellIn: 9},
				{Name: "Aged Brie", Quality: 12, SellIn: -1},
			},
		},
		{
			name: "Quality of an item Cannot increase past 50",
			args: []*Item{
				{Name: "Aged Brie", Quality: 50, SellIn: 10},
				{Name: "Aged Brie", Quality: 49, SellIn: 0},
			},
			expected: []*Item{
				{Name: "Aged Brie", Quality: 50, SellIn: 9},
				{Name: "Aged Brie", Quality: 50, SellIn: -1},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateQuality(tt.args)

			for idx, got := range tt.args {
				want := tt.expected[idx]

				if !reflect.DeepEqual(got, want) {
					t.Errorf("got: %+v  want:%+v", got, want)
				}
			}
		})
	}
}
func Test_LegendaryItems_UpdateQuality_Rules(t *testing.T) {
	tests := []struct {
		name     string
		args     []*Item
		expected []*Item
	}{
		{
			name: "Legendary Items never have to be sold, and their Quality is always 80",
			args: []*Item{
				{Name: "Sulfuras, Hand of Ragnaros", Quality: 80, SellIn: 10},
				{Name: "Sulfuras, Hand of Ragnaros", Quality: 80, SellIn: 0},
				{Name: "Sulfuras, Hand of Ragnaros", Quality: 80, SellIn: -10},
			},
			expected: []*Item{
				{Name: "Sulfuras, Hand of Ragnaros", Quality: 80, SellIn: 10},
				{Name: "Sulfuras, Hand of Ragnaros", Quality: 80, SellIn: 0},
				{Name: "Sulfuras, Hand of Ragnaros", Quality: 80, SellIn: -10},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateQuality(tt.args)

			for idx, got := range tt.args {
				want := tt.expected[idx]

				if !reflect.DeepEqual(got, want) {
					t.Errorf("got: %+v  want:%+v", got, want)
				}
			}
		})
	}
}
func Test_BackstagePasses_UpdateQuality_Rules(t *testing.T) {
	tests := []struct {
		name     string
		args     []*Item
		expected []*Item
	}{
		{
			name: "Backstage passes increases by 2 when SellIn is between 10 and 6 (inclusive)",
			args: []*Item{
				{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: 0, SellIn: 11},
				{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: 0, SellIn: 10},
				{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: 0, SellIn: 6},
			},
			expected: []*Item{
				{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: 1, SellIn: 10},
				{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: 2, SellIn: 9},
				{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: 2, SellIn: 5},
			},
		},
		{
			name: "Backstage passes increases by 3 when SellIn is between 5 and 1 (inclusive)",
			args: []*Item{
				{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: 0, SellIn: 5},
				{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: 0, SellIn: 1},
			},
			expected: []*Item{
				{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: 3, SellIn: 4},
				{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: 3, SellIn: 0},
			},
		},
		{
			name: "Backstage passes Quality drops to 0 when SellIn is 0 or less",
			args: []*Item{
				{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: 50, SellIn: 0},
				{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: 0, SellIn: -1},
			},
			expected: []*Item{
				{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: 0, SellIn: -1},
				{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: 0, SellIn: -2},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateQuality(tt.args)

			for idx, got := range tt.args {
				want := tt.expected[idx]

				if !reflect.DeepEqual(got, want) {
					t.Errorf("got: %+v  want:%+v", got, want)
				}
			}
		})
	}
}

//endregion -- Regression Tests

// region -- New Functionality Tests

func Test_Conjured_UpdateQuality_Rules(t *testing.T) {
	tests := []struct {
		name     string
		args     []*Item
		expected []*Item
	}{
		{
			name: "Conjured items degrade in Quality twice as fast as normal items",
			args: []*Item{
				{Name: "Rabbit's Foot", Quality: 1, SellIn: 1},
				{Name: "Conjured Mana Cake", Quality: 2, SellIn: 1},
				{Name: "Rabbit's Foot", Quality: 2, SellIn: 0},
				{Name: "Conjured Mana Cake", Quality: 4, SellIn: 0},
			},
			expected: []*Item{
				{Name: "Rabbit's Foot", Quality: 0, SellIn: 0},
				{Name: "Conjured Mana Cake", Quality: 0, SellIn: 0},
				{Name: "Rabbit's Foot", Quality: 0, SellIn: -1},
				{Name: "Conjured Mana Cake", Quality: 0, SellIn: -1},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateQuality(tt.args)

			for idx, got := range tt.args {
				want := tt.expected[idx]

				if !reflect.DeepEqual(got, want) {
					t.Errorf("got: %+v  want:%+v", got, want)
				}
			}
		})
	}
}

//endregion -- New Functionality Tests
