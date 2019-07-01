package cointop

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/miguelmota/cointop/cointop/common/pad"
)

func (ct *Cointop) togglePortfolio() error {
	ct.State.filterByFavorites = false
	ct.State.portfolioVisible = !ct.State.portfolioVisible
	go ct.updateTable()
	return nil
}

func (ct *Cointop) toggleShowPortfolio() error {
	ct.State.filterByFavorites = false
	ct.State.portfolioVisible = true
	go ct.updateTable()
	return nil
}

func (ct *Cointop) togglePortfolioUpdateMenu() error {
	ct.State.portfolioUpdateMenuVisible = !ct.State.portfolioUpdateMenuVisible
	if ct.State.portfolioUpdateMenuVisible {
		return ct.showPortfolioUpdateMenu()
	}
	return ct.hidePortfolioUpdateMenu()
}

func (ct *Cointop) updatePortfolioUpdateMenu() {
	coin := ct.highlightedRowCoin()
	exists := ct.PortfolioEntryExists(coin)
	value := strconv.FormatFloat(coin.Holdings, 'f', -1, 64)
	var mode string
	var current string
	var submitText string
	if exists {
		mode = "Edit"
		current = fmt.Sprintf("(current %s %s)", value, coin.Symbol)
		submitText = "Set"
	} else {
		mode = "Add"
		submitText = "Add"
	}
	header := ct.colorscheme.MenuHeader(fmt.Sprintf(" %s Portfolio Entry %s\n\n", mode, pad.Left("[q] close ", ct.maxTableWidth-26, " ")))
	label := fmt.Sprintf(" Enter holdings for %s %s", ct.colorscheme.MenuLabel(coin.Name), current)
	content := fmt.Sprintf("%s\n%s\n\n%s%s\n\n\n [Enter] %s    [ESC] Cancel", header, label, strings.Repeat(" ", 29), coin.Symbol, submitText)

	ct.update(func() {
		ct.Views.PortfolioUpdateMenu.Backing.Clear()
		ct.Views.PortfolioUpdateMenu.Backing.Frame = true
		fmt.Fprintln(ct.Views.PortfolioUpdateMenu.Backing, content)
		fmt.Fprintln(ct.Views.Input.Backing, value)
		ct.Views.Input.Backing.SetCursor(len(value), 0)
	})
}

func (ct *Cointop) showPortfolioUpdateMenu() error {
	coin := ct.highlightedRowCoin()
	if coin == nil {
		ct.togglePortfolio()
		return nil
	}

	ct.State.portfolioUpdateMenuVisible = true
	ct.updatePortfolioUpdateMenu()
	ct.setActiveView(ct.Views.PortfolioUpdateMenu.Name)
	return nil
}

func (ct *Cointop) hidePortfolioUpdateMenu() error {
	ct.State.portfolioUpdateMenuVisible = false
	ct.setViewOnBottom(ct.Views.PortfolioUpdateMenu.Name)
	ct.setViewOnBottom(ct.Views.Input.Name)
	ct.setActiveView(ct.Views.Table.Name)
	ct.update(func() {
		if ct.Views.PortfolioUpdateMenu.Backing == nil {
			return
		}

		ct.Views.PortfolioUpdateMenu.Backing.Clear()
		ct.Views.PortfolioUpdateMenu.Backing.Frame = false
		fmt.Fprintln(ct.Views.PortfolioUpdateMenu.Backing, "")

		ct.Views.Input.Backing.Clear()
		fmt.Fprintln(ct.Views.Input.Backing, "")
	})
	return nil
}

// sets portfolio entry holdings from inputed value
func (ct *Cointop) setPortfolioHoldings() error {
	defer ct.hidePortfolioUpdateMenu()
	coin := ct.highlightedRowCoin()

	b := make([]byte, 100)
	n, err := ct.Views.Input.Backing.Read(b)
	if n == 0 {
		return nil
	}

	value := normalizeFloatstring(string(b))
	shouldDelete := value == ""
	var holdings float64

	if !shouldDelete {
		holdings, err = strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
	}

	ct.setPortfolioEntry(coin.Name, holdings)

	if shouldDelete {
		ct.removePortfolioEntry(coin.Name)
		ct.updateTable()
		ct.goToGlobalIndex(0)
	} else {
		ct.updateTable()
		ct.goToGlobalIndex(coin.Rank - 1)
	}

	return nil
}

func (ct *Cointop) PortfolioEntry(c *Coin) (*PortfolioEntry, bool) {
	if c == nil {
		return &PortfolioEntry{}, true
	}

	var p *PortfolioEntry
	var isNew bool
	var ok bool
	key := strings.ToLower(c.Name)
	if p, ok = ct.State.portfolio.Entries[key]; !ok {
		// NOTE: if not found then try the symbol
		key := strings.ToLower(c.Symbol)
		if p, ok = ct.State.portfolio.Entries[key]; !ok {
			p = &PortfolioEntry{
				Coin:     c.Name,
				Holdings: 0,
			}
			isNew = true
		}
	}

	return p, isNew
}

func (ct *Cointop) setPortfolioEntry(coin string, holdings float64) {
	c, _ := ct.State.allCoinsSlugMap[strings.ToLower(coin)]
	p, isNew := ct.PortfolioEntry(c)
	if isNew {
		key := strings.ToLower(coin)
		ct.State.portfolio.Entries[key] = &PortfolioEntry{
			Coin:     coin,
			Holdings: holdings,
		}
	} else {
		p.Holdings = holdings
	}
}

func (ct *Cointop) removePortfolioEntry(coin string) {
	delete(ct.State.portfolio.Entries, strings.ToLower(coin))
}

func (ct *Cointop) PortfolioEntryExists(c *Coin) bool {
	_, isNew := ct.PortfolioEntry(c)
	return !isNew
}

func (ct *Cointop) portfolioEntriesCount() int {
	return len(ct.State.portfolio.Entries)
}

func (ct *Cointop) getPortfolioSlice() []*Coin {
	sliced := []*Coin{}
	for i := range ct.State.allCoins {
		if ct.portfolioEntriesCount() == 0 {
			break
		}
		coin := ct.State.allCoins[i]
		p, isNew := ct.PortfolioEntry(coin)
		if isNew {
			continue
		}
		coin.Holdings = p.Holdings
		balance := coin.Price * p.Holdings
		balancestr := fmt.Sprintf("%.2f", balance)
		if ct.State.currencyConversion == "ETH" || ct.State.currencyConversion == "BTC" {
			balancestr = fmt.Sprintf("%.5f", balance)
		}
		balance, _ = strconv.ParseFloat(balancestr, 64)
		coin.Balance = balance
		sliced = append(sliced, coin)
	}

	sort.Slice(sliced, func(i, j int) bool {
		return sliced[i].Balance > sliced[j].Balance
	})

	for i, coin := range sliced {
		coin.Rank = i + 1
	}

	return sliced
}

func (ct *Cointop) getPortfolioTotal() float64 {
	portfolio := ct.getPortfolioSlice()
	var total float64
	for _, p := range portfolio {
		total += p.Balance
	}
	return total
}

func normalizeFloatstring(input string) string {
	re := regexp.MustCompile(`(\d+\.\d+|\.\d+|\d+)`)
	result := re.FindStringSubmatch(input)
	if len(result) > 0 {
		return result[0]
	}

	return ""
}
