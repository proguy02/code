package main

import (
	"code/utils"
	"fmt"
)

func main() {
	fmt.Println(utils.Base("Simply add 42 (hex) and 10 (bin) and you will see the result is 68."))
	fmt.Println(utils.String("it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."))
}
