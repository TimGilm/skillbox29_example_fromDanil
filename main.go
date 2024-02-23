package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	in := scanner(&wg)
	middle := summator(in, &wg)
	out := multiplier(middle, &wg)
	receiver(out, &wg)

	wg.Wait()
	fmt.Println("Конец!!!")
}
func scanner(wg *sync.WaitGroup) chan int {
	out := make(chan int)
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("Отправитель завершает работу")
			wg.Done()
		}()
		defer func() {
			fmt.Println("отправитель закрывает канал")
			close(out)
		}()
		var scan string
		var digit int
		for {
			_, err := fmt.Scan(&scan)
			if err != nil {
				log.Println(err)
				continue
			}
			digit, err = strconv.Atoi(scan)
			if err != nil {
				if scan == "стоп" {
					break
				}
				log.Println(err)
				continue
			}
			fmt.Printf("Отправитель отправил %v\n", digit)
			out <- digit
		}
	}()
	return out
}
func summator(in chan int, wg *sync.WaitGroup) chan int {
	out := make(chan int, 5)
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("Сумматор завершает работу")
			wg.Done()
		}()
		defer func() {
			fmt.Println("Сумматор закрывает канал")
			close(out)
		}()
		for value := range in {
			result := value + value
			fmt.Println("Сумматор принял %v, отправил %v\n", value, result)
			out <- result
		}
	}()
	return out
}
func multiplier(in chan int, wg *sync.WaitGroup) chan int {
	out := make(chan int, 5)
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("Мультипликатор завершает работу")
			wg.Done()
		}()
		defer func() {
			fmt.Println("Мультипликатор закрывает канал")
			close(out)
		}()
		for value := range in {
			result := value * value
			fmt.Printf("Мульпликатор принял %v, отправил %v\n", value, result)
			out <- result
		}
	}()
	return out
}
func receiver(in chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Printf("Получатель завершает работу")
			wg.Done()
		}()
		for value := range in {
			fmt.Printf("Получатель принял %v\n", value)
		}
	}()
}
