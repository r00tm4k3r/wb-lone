package main

import (
	"fmt"
	"sync"
)

func concurrentPow(v int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done() // сообщает wait group о том, чтобы выкинуть данную функцию из очереди ожидания
	//(ставится при завершении горутины)
	ch <- v * v // записываем в канал наш результат
}

func main() {
	vals := []int{2, 4, 6, 8, 10} //создаем массив чисел, которые указаны в задании
	ch := make(chan int)          // созаем канал, где будут храниться наши выходные данные
	wg := new(sync.WaitGroup)
	// ^ создаем wait group для того, чтобы синхронизировать наши потоки и чтобы программа не
	// не завершилась раньше времени при условии того, что не все потоки еще отработали свое

	for i := 0; i < 5; i++ { //цикл для прогонки массива
		wg.Add(1)                         // говорим wait group, что нужно подождать 1 элемент
		go concurrentPow(vals[i], wg, ch) // запускаем функцию возведения в степень горутиной
	}

	go func() { // функция которая блокирует wait group и закрывает канал
		// запускаестя в горутине, потому что при последовательном закрытии у нас получается dealock канала
		wg.Wait() // ждем пока наш счетчик на wait group будет ровняться 0
		close(ch) // после закрываем канал
	}()

	for val := range ch { // запускаем цикл прохождения по каналу
		fmt.Println(val) // выводим результат из канала
	}

}
