package client

import (
	"fmt"
)

func Foo() {
	fmt.Println("hello", Channel{
		ChannelID: []byte{0},
	})
	// fmt.Println("foo")
}
