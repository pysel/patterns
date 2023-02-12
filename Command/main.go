/*
Type: behavioral
Command is a software desing pattern used to decouple invoking of specific functionality on a reciever (the object affected by some signal)
and client (sender of a signal).

An example would be a remote control, which is able to turn on the music or lights (one at a time!!!)

The idea is for a remote control object to contain one `command` object, which should have `.exec()` method.
When an improvised button is pressed on a remote, it does not care which object it currently sends signal to.
It can just run a `exec` command at the reciever object.
*/
package main

type Command interface {
	exec()
}

type remote struct {
	cmd Command
}

func (r remote) pressButton() {
	r.cmd.exec()
}

func main() {
	// configure remote to turn lights on when pressing button
	remote := remote{
		cmd: LightOnCommand{Light: Light{}},
	}
	remote.pressButton()

	// configure remote to turn music on when pressing button
	remote.cmd = MusicOnCommand{
		Music: Music{},
	}
	remote.pressButton()

	// configure remote to turn lights off when pressing button
	remote.cmd = LightOffCommand{
		Light: Light{},
	}
	remote.pressButton()

	// music still on, enjoy :)
}
