package cli

import "context"

func (this *CLI) deleteCharacter() {
	if len(this.split) != 3 {
		this.sendMessage("command usage: character delete <name>")
		return
	}

	characterName := this.split[2]
	conn, _ := this.Manager.DB.Acquire(context.TODO())
	defer conn.Release()

	const query = "DELETE FROM characters WHERE name=$1;"
	t, err := conn.Exec(context.TODO(), query, characterName)
	if t.RowsAffected() == 0 || err != nil {
		this.sendMessage("error: character not found")
		return
	}
	this.sendMessage("character deleted")
}

func (this *CLI) character() {
	m := map[string]func() {
		"delete": this.deleteCharacter,
	}

	if f := this.getFunc(m, 1); f != nil {
		f()
	}
}
