package utils

import "fmt"

func FormatTime(second uint64) string {
	minutes := second / 60
	hours := second / 3600
	days := second / 86400
	return fmt.Sprintf("%d days, %d hours, %d minutes",
		days,
		hours%24,
		minutes%60,
	)
}
