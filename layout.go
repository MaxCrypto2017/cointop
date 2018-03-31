package main

import (
	"github.com/jroimartin/gocui"
	"github.com/miguelmota/cointop/pkg/pad"
	"github.com/miguelmota/cointop/pkg/table"
)

func (ct *Cointop) layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	chartHeight := 10
	topOffset := 0

	if v, err := g.SetView("market", 0, topOffset, maxX, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		ct.marketview = v
		ct.marketview.Frame = false
		ct.marketview.BgColor = gocui.ColorBlack
		ct.marketview.FgColor = gocui.ColorWhite
		ct.updateMarket()
	}

	topOffset = topOffset + 1
	if v, err := g.SetView("chart", 0, topOffset, maxX, topOffset+chartHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		ct.chartview = v
		ct.chartview.Frame = false
		ct.updateChart()
	}

	topOffset = topOffset + chartHeight
	if v, err := g.SetView("header", 0, topOffset, maxX, topOffset+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		t := table.New().SetWidth(maxX)
		headers := []string{
			pad.Right("[r]ank", 13, " "),
			pad.Right("[n]ame", 13, " "),
			pad.Right("[s]ymbol", 8, " "),
			pad.Left("[p]rice", 10, " "),
			pad.Left("[m]arket cap", 17, " "),
			pad.Left("24H [v]olume", 15, " "),
			pad.Left("[1]H%", 9, " "),
			pad.Left("[2]4H%", 9, " "),
			pad.Left("[7]D%", 9, " "),
			pad.Left("[t]otal supply", 20, " "),
			pad.Left("[a]vailable supply", 19, " "),
			pad.Left("[l]ast updated", 17, " "),
		}
		for _, h := range headers {
			t.AddCol(h)
		}

		t.Format().Fprint(v)
		ct.headersview = v
		ct.headersview.Frame = false
		ct.headersview.Highlight = true
		ct.headersview.SelBgColor = gocui.ColorGreen
		ct.headersview.SelFgColor = gocui.ColorBlack
	}

	topOffset = topOffset + 1
	if v, err := g.SetView("table", 0, topOffset, maxX, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		ct.tableview = v
		ct.tableview.Frame = false
		ct.tableview.Highlight = true
		ct.tableview.SelBgColor = gocui.ColorCyan
		ct.tableview.SelFgColor = gocui.ColorBlack
		ct.updateTable()
		ct.rowChanged()
	}

	if v, err := g.SetView("status", 0, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		ct.statusview = v
		ct.statusview.Frame = false
		ct.statusview.BgColor = gocui.ColorCyan
		ct.statusview.FgColor = gocui.ColorBlack
		ct.updateStatus("")
	}

	return nil
}
