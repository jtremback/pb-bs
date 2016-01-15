package client

import (
	"fmt"
)

func Foo() {
	fmt.Println("hello", Channel{
		ChannelID: []byte{4, 20},
	})
	// fmt.Println("foo")
}
