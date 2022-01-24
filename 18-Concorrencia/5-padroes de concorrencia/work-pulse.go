package main

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func worker(tarefas chan int, resultados chan int) {

}
func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

}
