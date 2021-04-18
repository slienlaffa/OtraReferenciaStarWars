package communications

import (
	"math"
	"strings"
)

type Coordinates struct {
	x, y float32
}

var satellites = map[string]Coordinates{
	"Kenobi":    {-500, -200},
	"Skywalker": {100, -100},
	"Sato":      {500, 100},
}

// input: distancia al emisor tal cual se recibe en cada satélite
// output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
func GetLocation(distances ...float32) (x, y float32) {
	r1, r2, r3 := distances[0], distances[1], distances[2]
	A := 2*satellites["Skywalker"].x - 2*satellites["Kenobi"].x
	B := 2*satellites["Skywalker"].y - 2*satellites["Kenobi"].y
	C := pow2(r1) - pow2(r2) - pow2(satellites["Kenobi"].x) + pow2(satellites["Skywalker"].x) - pow2(satellites["Kenobi"].y) + pow2(satellites["Skywalker"].y)
	D := 2*satellites["Sato"].x - 2*satellites["Skywalker"].x
	E := 2*satellites["Sato"].y - 2*satellites["Skywalker"].y
	F := pow2(r2) - pow2(r3) - pow2(satellites["Skywalker"].x) + pow2(satellites["Sato"].x) - pow2(satellites["Skywalker"].y) + pow2(satellites["Sato"].y)
	emisorX := (C*E - F*B) / (E*A - B*D)
	emisorY := (C*D - A*F) / (B*D - A*E)
	return round(emisorX), round(emisorY)
}

// redondea un float para que solo tenga 2 decimales
func round(number float32) float32 {
	return float32(math.Round(float64(number)*100) / 100)
}

// math.Pow solo acepta float64, asi que para evitar demasiadas conversiones
func pow2(number float32) float32 {
	return number * number
}

// input: el mensaje tal cual es recibido en cada satélite
// output: el mensaje tal cual lo genera el emisor del mensaje
func GetMessage(messages ...[]string) (msg string) {
	messages = limpiarDesfase(messages)
	completeMessage := messages[0]
	found := false

	for i := 1; i < len(messages); i++ {
		for index, word := range messages[i] {
			for _, inMessage := range completeMessage {
				if word == inMessage {
					found = true
					break
				}
			}
			if !found && completeMessage[index] == "" {
				completeMessage[index] = word
			}
			found = false
		}
	}
	return strings.Join(completeMessage, " ")
}

// estoy asumiendo que desfase son espacios blancos solo en el principio
func limpiarDesfase(messages [][]string) [][]string {
	lenght := len(messages[0])
	shorter := false
	for {
		for _, message := range messages {
			if lenght != len(message) {
				if lenght > len(message) {
					lenght = len(message)
				}
				shorter = true
				break
			}
		}
		if !shorter {
			break
		}
		if shorter {
			for i, message := range messages {
				if lenght < len(message) {
					messages[i] = message[1:]
				} else {
					messages[i] = message
				}
			}
			shorter = false
		}
	}
	return messages
}
