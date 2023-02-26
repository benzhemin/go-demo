package main

import StringUtil "go-lang/morestrings"
import "fmt"

func main() {
	var s = "hello world"

	var reserve = StringUtil.ReverseRunes(s);

	fmt.Println(reserve);
}
