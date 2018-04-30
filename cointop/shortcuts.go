package cointop

func defaultShortcuts() map[string]string {
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
		"esc":       "quit",
		"space":     "toggle_favorite",
		"ctrl+c":    "quit",
		"ctrl+d":    "page_down",
		"ctrl+f":    "open_search",
		"ctrl+n":    "next_page",
		"ctrl+p":    "previous_page",
		"ctrl+r":    "refresh",
		"ctrl+s":    "save",
		"ctrl+u":    "page_up",
		"alt+up":    "sort_column_asc",
		"alt+down":  "sort_column_desc",
		"alt+left":  "sort_left_column",
		"alt+right": "sort_right_column",
		"F1":        "help",
		"F5":        "refresh",
		"0":         "first_page",
		"1":         "sort_column_1h_change",
		"2":         "sort_column_24h_change",
		"7":         "sort_column_7d_change",
		"a":         "sort_column_available_supply",
		"c":         "toggle_row_chart",
		"f":         "toggle_show_favorites",
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
		"p":         "sort_column_price",
		"r":         "sort_column_rank",
		"s":         "sort_column_symbol",
		"t":         "sort_column_total_supply",
		"u":         "sort_column_last_updated",
		"v":         "sort_column_24h_volume",
		"q":         "quit",
		"$":         "last_page",
		"?":         "help",
		"/":         "open_search",
	}
}
