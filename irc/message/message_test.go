package twitchirc

import "testing"

func TestParseLine(t *testing.T) {
	{
		got := ParseLine("")
		want := true
		if !got.isNil {
			t.Errorf("Want %t got %t", want, got.isNil)
		}
	}

	{
		got := ParseLine(":tmi.twitch.tv SOME RANDOM TEXT")
		want := true
		if !got.isNil {
			t.Errorf("Want %t got %t", want, got.isNil)
		}
	}

	{
		got := ParseLine("PING :tmi.twitch.tv")
		want := true
		if !got.isPing {
			t.Errorf("Want %t got %t", want, got.isNil)
		}
	}

	{
		got := ParseLine(":ricardinst!ricardinst@ricardinst.tmi.twitch.tv PRIVMSG #rafiusky :na real")
		want := Message{
			Sender:  "ricardinst",
			Text:    "na real",
			Channel: "rafiusky",
		}
		if got.Sender != want.Sender {
			t.Errorf("Want %s got %s", want.Sender, got.Sender)
		}
		if got.Text != want.Text {
			t.Errorf("Want %s got %s", want.Text, got.Text)
		}
		if got.Channel != want.Channel {
			t.Errorf("Want %s got %s", want.Channel, got.Channel)
		}
	}

}
