package main

import (
	"fmt"
	"testing"
)

func TestAddJoke(t *testing.T) {
	fmt.Println("TestAddJoke start:")
	err := AddJoke("喝水了！喝水了！喝水了！举起旁边的杯子，没有水就赶紧满上，感情深，一口闷！")
	err = AddJoke("喝水了！喝水了！喝水了！拿起旁边的水杯，Cheers！")
	err = AddJoke("美好的一天从喝水开始，你喝水了吗？")
	if err != nil {
		t.Errorf("TestDeleteUser error: err=%s", err)
	}
}

func TestSelectJokeList(t *testing.T) {
	fmt.Println("TestSelectJokeList start:")
	contents, err := SelectJokeList()
	if err != nil {
		t.Errorf("TestSelectJokeList error: err=%s", err)
	}
	fmt.Println("contents=", contents)
}
