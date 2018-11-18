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
		if err != nil {
			return n, err
		}

		total += n
		time.Sleep(time.Millisecond * time.Duration(50+rand.Intn(25)))
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

	stdin := bufio.NewReader(os.Stdin)

	printf("GREETINGS PROFESSOR FALKEN\n\n")
	stdin.ReadLine()
	printf("\nA STRANGE GAME.\n")
	printf("THE ONLY WINNING MOVE\n")
	printf("IS NOT TO PLAY.\n")
}
