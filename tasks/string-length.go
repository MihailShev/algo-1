package tasks

import "strconv"

type StringLength struct {

}

func (s StringLength) Run(data []string) string {
	if len(data) > 0 {
		str := []rune(data[0])
		return strconv.Itoa(len(str))
	}

	return ""
}