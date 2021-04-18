package communications

import (
	"testing"
)

func TestGetLocation(t *testing.T) {
	cases := []struct {
		inDis1 float32
		inDis2 float32
		inDis3 float32
		wantX  float32
		wantY  float32
	}{
		{670.82039325, 200, 400, 100, 100},
		{447.2135955, 500.0, 806.22577483, -300, 200},
	}
	for _, c := range cases {
		gotX, gotY := GetLocation(c.inDis1, c.inDis2, c.inDis3)
		if gotX != c.wantX && gotY != c.wantY {
			t.Errorf("GetLocation(%.2f, %.2f, %.2f) == %f, %f, requiere %.2f, %.2f", c.inDis1, c.inDis2, c.inDis3, gotX, gotY, c.wantX, c.wantY)
		}
	}
}

func TestGetMessage(t *testing.T) {
	cases := []struct {
		messages        [][]string
		completeMessage string
	}{
		{
			[][]string{
				{"este", "", "", "mensaje", ""},
				{"", "es", "", "", "secreto"},
				{"este", "", "un", "", ""},
			},
			"este es un mensaje secreto",
		},
		{
			[][]string{
				{"", "este", "", "", "mensaje"},
				{"", "es", "", ""},
				{"", "este", "", "un", ""},
			},
			"este es un mensaje",
		},
		{
			[][]string{
				{"", "", "este", "", "", ""},
				{"", "es", "", "email"},
				{"", "este", "", "un", ""},
			},
			"este es un email",
		},
	}
	for _, c := range cases {
		message := GetMessage(c.messages...)
		if message != c.completeMessage {
			t.Errorf("GetMessage() == %s, requiere %s", message, c.completeMessage)
		}
	}
}
