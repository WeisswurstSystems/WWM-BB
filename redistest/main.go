package main

import (
	"fmt"

	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/redis"
)

func main() {
	c := redis.NewStore("redis://weisswurstsystems:geheim@geheim:geheim")
	defer c.Close()

	c.Save(meeting.Meeting{ID: "0", Buyer: "xxx"})
	c.Save(meeting.Meeting{ID: "2", Buyer: "xxx"})

	fmt.Println(c.Find("0"))
	fmt.Println(c.Find("0"))

	fmt.Println(c.Find("2"))
	fmt.Println(c.FindAll())
	c.Delete("0")
}
