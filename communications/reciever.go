package communications

// input: distancia al emisor tal cual se recibe en cada satélite
// output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
func GetLocation(distances ...float32) (x, y float32)

// input: el mensaje tal cual es recibido en cada satélite
// output: el mensaje tal cual lo genera el emisor del mensaje
func GetMessage(messages ...[]string) (msg string)
