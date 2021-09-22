package twitchirc

import "testing"

func TestParseLine(t *testing.T) {
	{
		got := parseLine("")
		want := true
		if !got.isNil {
			t.Errorf("Want %t got %t", want, got.isNil)
		}
	}

	{
		got := parseLine(":tmi.twitch.tv SOME RANDOM TEXT")
		want := true
		if !got.isNil {
			t.Errorf("Want %t got %t", want, got.isNil)
		}
	}

	{
		got := parseLine("PING :tmi.twitch.tv")
		want := true
		if !got.isPing {
			t.Errorf("Want %t got %t", want, got.isNil)
		}
	}

	{
		got := parseLine(":ricardinst!ricardinst@ricardinst.tmi.twitch.tv PRIVMSG #rafiusky :na real")
		want := Message{
			Sender: "ricardinst",
			Text:   "na real",
		}
		if got.Sender != want.Sender {
			t.Errorf("Want %s got %s", want.Sender, got.Sender)
		}
		if got.Text != want.Text {
			t.Errorf("Want %s got %s", want.Text, got.Text)
		}
	}

}
