package cointop

// DefaultShortcuts is a map of the default shortcuts
func DefaultShortcuts() map[string]string {
	return map[string]string{
		"up":        "move_up",
		"down":      "move_down",
		"left":      "previous_page",
		"right":     "next_page",
		"pagedown":  "page_down",
		"pageup":    "page_up",
		"home":      "move_to_page_first_row",
		"end":       "move_to_page_last_row",
		"enter":     "toggle_row_chart",
		"esc":       "quit_view",
		"space":     "toggle_favorite",
		"tab":       "move_down_or_next_page",
		"ctrl+c":    "quit",
		"ctrl+C":    "quit",
		"ctrl+d":    "page_down",
		"ctrl+f":    "open_search",
		"ctrl+n":    "next_page",
		"ctrl+p":    "previous_page",
		"ctrl+r":    "refresh",
		"ctrl+R":    "refresh",
		"ctrl+s":    "save",
		"ctrl+S":    "save",
		"ctrl+u":    "page_up",
		"ctrl+j":    "enlarge_chart",
		"ctrl+k":    "shorten_chart",
		"|":         "toggle_chart_fullscreen",
		"alt+up":    "sort_column_asc",
		"alt+down":  "sort_column_desc",
		"alt+left":  "sort_left_column",
		"alt+right": "sort_right_column",
		"F1":        "help",
		"F5":        "refresh",
		"0":         "first_page",
		"1":         "sort_column_1h_change",
		"2":         "sort_column_24h_change",
		"3":         "sort_column_30d_change",
		"6":         "sort_column_1y_change",
		"7":         "sort_column_7d_change",
		"a":         "sort_column_available_supply",
		"b":         "sort_column_balance",
		"c":         "show_currency_convert_menu",
		"C":         "show_currency_convert_menu",
		"e":         "show_portfolio_edit_menu",
		"E":         "show_portfolio_edit_menu",
		"A":         "toggle_price_alerts",
		"f":         "toggle_favorite",
		"F":         "toggle_show_favorites",
		"g":         "move_to_page_first_row",
		"G":         "move_to_page_last_row",
		"h":         "previous_page",
		"H":         "move_to_page_visible_first_row",
		"j":         "move_down",
		"k":         "move_up",
		"l":         "next_page",
		"L":         "move_to_page_visible_last_row",
		"m":         "sort_column_market_cap",
		"M":         "move_to_page_visible_middle_row",
		"n":         "sort_column_name",
		"o":         "open_link",
		"O":         "open_link",
		"p":         "sort_column_price",
		"P":         "toggle_portfolio",
		"r":         "sort_column_rank",
		"s":         "sort_column_symbol",
		"t":         "sort_column_total_supply",
		"u":         "sort_column_last_updated",
		"v":         "sort_column_24h_volume",
		"q":         "quit_view",
		"Q":         "quit_view",
		"%":         "sort_column_percent_holdings",
		"$":         "last_page",
		"?":         "help",
		"/":         "open_search",
		"]":         "next_chart_range",
		"[":         "previous_chart_range",
		"}":         "last_chart_range",
		"{":         "first_chart_range",
		">":         "scroll_right",
		"<":         "scroll_left",
		"+":         "show_price_alert_add_menu",
		"\\\\":      "toggle_table_fullscreen",
	}
}
