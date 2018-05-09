package cointop

import (
	"sort"

	"github.com/miguelmota/cointop/pkg/gocui"
)

func (ct *Cointop) sort(sortby string, desc bool, list []*coin) {
	ct.sortby = sortby
	ct.sortdesc = desc
	sort.Slice(list[:], func(i, j int) bool {
		if ct.sortdesc {
			i, j = j, i
		}
		a := list[i]
		b := list[j]
		switch sortby {
		case "rank":
			return a.Rank < b.Rank
		case "name":
			return a.Name < b.Name
		case "symbol":
			return a.Symbol < b.Symbol
		case "price":
			return a.Price < b.Price
		case "marketcap":
			return a.MarketCap < b.MarketCap
		case "24hvolume":
			return a.Volume24H < b.Volume24H
		case "1hchange":
			return a.PercentChange1H < b.PercentChange1H
		case "24hchange":
			return a.PercentChange24H < b.PercentChange24H
		case "7dchange":
			return a.PercentChange7D < b.PercentChange7D
		case "totalsupply":
			return a.TotalSupply < b.TotalSupply
		case "availablesupply":
			return a.AvailableSupply < b.AvailableSupply
		case "lastupdated":
			return a.LastUpdated < b.LastUpdated
		default:
			return a.Rank < b.Rank
		}
	})
	ct.updateHeaders()
}

func (ct *Cointop) sortfn(sortby string, desc bool) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if ct.sortby == sortby {
			desc = !desc
		}

		ct.sort(sortby, desc, ct.coins)
		ct.updateTable()
		return nil
	}
}

func (ct *Cointop) getSortColIndex() int {
	for i, col := range ct.tablecolumnorder {
		if ct.sortby == col {
			return i
		}
	}
	return 0
}

func (ct *Cointop) sortAsc() error {
	ct.sortdesc = false
	ct.sort(ct.sortby, ct.sortdesc, ct.coins)
	ct.updateTable()
	return nil
}

func (ct *Cointop) sortDesc() error {
	ct.sortdesc = true
	ct.sort(ct.sortby, ct.sortdesc, ct.coins)
	ct.updateTable()
	return nil
}

func (ct *Cointop) sortPrevCol() error {
	nextsortby := ct.tablecolumnorder[0]
	i := ct.getSortColIndex()
	k := i - 1
	if k < 0 {
		k = 0
	}
	nextsortby = ct.tablecolumnorder[k]
	ct.sort(nextsortby, ct.sortdesc, ct.coins)
	ct.updateTable()
	return nil
}

func (ct *Cointop) sortNextCol() error {
	nextsortby := ct.tablecolumnorder[0]
	l := len(ct.tablecolumnorder)
	i := ct.getSortColIndex()
	k := i + 1
	if k > l-1 {
		k = l - 1
	}
	nextsortby = ct.tablecolumnorder[k]
	ct.sort(nextsortby, ct.sortdesc, ct.coins)
	ct.updateTable()
	return nil
}
