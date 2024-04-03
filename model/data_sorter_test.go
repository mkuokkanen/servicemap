package model

import (
	"fmt"
	"github.com/mkuokkanen/servicemap/pkl/gen"
	"testing"
)

func TestSortData(t *testing.T) {
	d := gen.Data{
		Groups: []*gen.Group{
			{
				Name: "G1",
				Sort: 1,
				Services: []*gen.Service{
					{
						Name: "G1 S1",
						Components: []*gen.Component{
							{Name: "G1 S1 C1"},
							{Name: "G1 S1 C0"},
						}},
					{
						Name: "G1 S0",
						Components: []*gen.Component{
							{Name: "G1 S0 C1"},
							{Name: "G1 S0 C0"},
						}}}},
			{
				Name: "G0",
				Sort: 0,
				Services: []*gen.Service{
					{
						Name: "G0 S1",
						Components: []*gen.Component{
							{Name: "G0 S1 C1"},
							{Name: "G0 S1 C0"},
						}},
					{
						Name: "G0 S0",
						Components: []*gen.Component{
							{Name: "G0 S0 C1"},
							{Name: "G0 S0 C0"},
						}}}},
		},
	}

	sortData(&d)

	for g := range 2 {
		gStr := fmt.Sprintf("G%d", g)
		if d.Groups[g].Name != gStr {
			t.Fatalf("Expected " + gStr + " got " + d.Groups[g].Name)
		}

		for s := range 2 {
			sStr := fmt.Sprintf("G%d S%d", g, s)
			if d.Groups[g].Services[s].Name != sStr {
				t.Fatalf("Expected " + sStr + " got " + d.Groups[g].Services[s].Name)
			}

			for c := range 2 {
				cStr := fmt.Sprintf("G%d S%d C%d", g, s, c)
				if d.Groups[g].Services[s].Components[c].Name != cStr {
					t.Fatalf("Expected " + cStr + " got " + d.Groups[g].Services[s].Components[c].Name)
				}
			}
		}
	}
}
