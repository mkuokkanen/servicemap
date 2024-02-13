package model

import (
	"cmp"
	"slices"
	"strings"
)

func sortData(data *Data) {
	slices.SortFunc(data.Groups, func(a, b GroupData) int {
		return cmp.Compare(a.Sort, b.Sort)
	})
	for _, group := range data.Groups {
		sortServices(group.Services)
	}
}

func sortServices(products []ServiceData) {
	slices.SortFunc(products, func(a, b ServiceData) int {
		return cmp.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})
	for _, p := range products {
		sortComponents(p.Components)
	}
}

func sortComponents(components []ComponentData) {
	slices.SortFunc(components, func(a, b ComponentData) int {
		return cmp.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})
}
