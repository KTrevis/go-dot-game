package cli

import "strings"

func getArgumentList(m map[string]func()) string {
	str := "\n"

	for key := range m {
		str += "- " + key + "\n"
	}

	str = strings.TrimSuffix(str, "\n")
	return str
}

func (this *CLI) getFunc(m map[string]func(), index int) func() {
	if len(this.split) <= index {
		msg := "invalid argument. arguments possible: "
		msg += getArgumentList(m)
		this.sendMessage(msg)
		return nil
	}

	fn, ok := m[this.split[index]]
	if !ok {
		msg := "invalid command. commands possible: "
		msg += getArgumentList(m)
		this.sendMessage(msg)
		return nil
	}
	return fn
}
