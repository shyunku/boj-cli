package solvedac

import "fmt"

type Tier struct {
	Label string
	Color string // ANSI 256-color escape
	Reset string
}

var reset = "\033[0m"

// tierInfo maps level (0-30) to display info.
func TierInfo(level int) Tier {
	switch {
	case level == 0:
		return Tier{"??", "\033[38;5;242m", reset} // gray (unrated)
	case level <= 5:
		roman := romanTier(level)
		return Tier{fmt.Sprintf("B%s", roman), "\033[38;5;130m", reset} // brown (bronze)
	case level <= 10:
		roman := romanTier(level - 5)
		return Tier{fmt.Sprintf("S%s", roman), "\033[38;5;250m", reset} // light gray (silver)
	case level <= 15:
		roman := romanTier(level - 10)
		return Tier{fmt.Sprintf("G%s", roman), "\033[38;5;220m", reset} // yellow (gold)
	case level <= 20:
		roman := romanTier(level - 15)
		return Tier{fmt.Sprintf("P%s", roman), "\033[38;5;51m", reset} // cyan (platinum)
	case level <= 25:
		roman := romanTier(level - 20)
		return Tier{fmt.Sprintf("D%s", roman), "\033[38;5;45m", reset} // blue (diamond)
	case level <= 30:
		roman := romanTier(level - 25)
		return Tier{fmt.Sprintf("R%s", roman), "\033[38;5;205m", reset} // pink (ruby)
	default:
		return Tier{"??", "\033[38;5;242m", reset}
	}
}

// romanTier converts 1-5 to V-I (BOJ tiers go V=hardest within group... wait, V is lowest)
// In BOJ: Bronze V (easiest) → Bronze I (hardest within bronze)
// level 1 = Bronze V, level 5 = Bronze I
// So offset 1→"V", 2→"IV", 3→"III", 4→"II", 5→"I"
func romanTier(offset int) string {
	switch offset {
	case 1:
		return "5"
	case 2:
		return "4"
	case 3:
		return "3"
	case 4:
		return "2"
	case 5:
		return "1"
	default:
		return "?"
	}
}
