package cli

func (this *CLI) echo(split []string) {
	msg := ""

	for i, str := range split {
		if i == 0 {
			continue
		}
		msg += str + " "
	}
	this.sendMessage(msg)
}
