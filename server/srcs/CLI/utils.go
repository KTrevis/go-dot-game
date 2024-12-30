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

func (this *CLI) validArg(m map[string]func()) func() {
	if len(this.split) < 2 {
		msg := "invalid command. arguments possible: "
		msg += getArgumentList(m)
		this.sendMessage(msg)
		return nil
	}

	fn, ok := m[this.split[1]]
	if !ok {
		msg := "invalid argument. arguments possible: "
		msg += getArgumentList(m)
		this.sendMessage(msg)
		return nil
	}
	return fn
}
