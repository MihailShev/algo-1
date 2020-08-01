package tasks

import "strconv"

type LuckyTickets struct {

}

func (l LuckyTickets) Run(data []string) string {
	if len(data) > 0 {
		str := []rune(data[0])
		return strconv.Itoa(len(str))
	}

	return ""
}
