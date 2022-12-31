package goarea

import "math"

/*
para que seja um função exportada tem que ser realizado o comentario da seguinte forma {Nome da função } [ o que ela faz e o que ela realiza].
pode colocar somente o nome do pacote e não o caminho completo de para/sub-pasta e com isso colocar somente a pasta.
você coloca o nome do pacote e o nome do arquivo são os mesmos para ficar mais fácil de achar
*/

// Pi é uma proporção numérica definida
const Pi = 3.1415

// Circ é responsável por calcular a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsavel para calcular a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// não é visivel
func _Triangulo(base, altura float64) float64 {
	return (base * altura) / 2
}
