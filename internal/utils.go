package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ExpandTilde(p string) string {
	if strings.HasPrefix(p, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		return filepath.Join(home, p[2:])
	}
	return p
}

func HumanBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}

func RemoveElement(a []string, val string) []string {
	x := a[:0]
	for _, v := range a {
		if val != v {
			x = append(x, v)
		}
	}
	return x
}




func RemoveDup(a []string) []string {
	var n []string
	dups := make(map[string]bool)
	for _, v := range a {
		_, ok := dups[v]
		if !ok {
			dups[v] = true
			n = append(n, v)
		} else {
			continue
		}
	}
	return n
}
