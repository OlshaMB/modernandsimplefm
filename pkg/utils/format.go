package utils

import (
	"fmt"
	"strings"
	"time"
)

func FormatTime(t time.Time, prefix string) string {
    diff := time.Since(t)
    switch {
    case diff.Seconds() < 60:
        return prefix + " a few seconds ago"
    case diff.Minutes() < 60:
        return fmt.Sprintf("%s %d minutes ago", prefix, int(diff.Minutes()))
    case diff.Hours() < 24:
        return fmt.Sprintf("%s %d hours ago", prefix, int(diff.Hours()))
    case diff.Hours() < 24*30:
        return fmt.Sprintf("%s %d days ago", prefix, int(diff.Hours()/24))
    case diff.Hours() < 24*365:
        return fmt.Sprintf("%s %d months ago", prefix, int(diff.Hours()/(24*30)))
    default:
        return fmt.Sprintf("%s %d years ago", prefix, int(diff.Hours()/(24*365)))
    }
}
func StripPathSeparator(path string) string {
	return strings.Trim(path, "/")
}