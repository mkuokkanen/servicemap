package model

import (
	"cmp"
	"github.com/mkuokkanen/servicemap/pkl/gen"
	"slices"
	"strings"
)

func sortData(data *gen.Data) {
	slices.SortFunc(data.Groups, func(a, b *gen.Group) int {
		return cmp.Compare(a.Sort, b.Sort)
	})
	for _, group := range data.Groups {
		sortServices(group.Services)
	}
}

func sortServices(products []*gen.Service) {
	slices.SortFunc(products, func(a, b *gen.Service) int {
		return cmp.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})
	for _, p := range products {
		sortComponents(p.Components)
	}
}

func sortComponents(components []*gen.Component) {
	slices.SortFunc(components, func(a, b *gen.Component) int {
		return cmp.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})
}
