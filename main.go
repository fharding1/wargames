package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	wordlist = "ABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890"
	target   = "CPE1704TKS"
	delay    = 75
)

func printf(format string, a ...interface{}) (n int, err error) {
	out := fmt.Sprintf(format, a...)

	var total int
	for _, ch := range out {
		n, err := fmt.Printf("%c", ch)
		total += n

		if err != nil {
			return total, err
		}

		time.Sleep(time.Millisecond * time.Duration(delay))
	}

	return total, nil
}

func main() {
	rand.Seed(time.Now().Unix())

	cur := make([]byte, len(target))
	wg := new(sync.WaitGroup)

	for i := 0; i < len(cur); i++ {
		wg.Add(1)
		go func(i int) {
			for cur[i] != target[i] {
				time.Sleep(time.Millisecond * delay)
				cur[i] = wordlist[rand.Intn(len(wordlist))]
				fmt.Printf("%s\r", cur)
			}

			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Printf("%s\n\n", cur)

	printf("GREETINGS PROFESSOR FALKEN\n\n")
	bufio.NewReader(os.Stdin).ReadLine()
	printf("\nA STRANGE GAME.\n")
	printf("THE ONLY WINNING MOVE\n")
	printf("IS NOT TO PLAY.\n")
}
