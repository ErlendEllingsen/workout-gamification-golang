package main

import "fmt"

type Badge struct {
	id   uint32
	name string
	desc string
}

var Badges = []Badge{
	Badge{
		id:   0x1,
		name: "dfe",
		desc: "gjjk",
	},
	Badge{
		id:   0x2,
		name: "asd",
		desc: "desc",
	},
	Badge{
		id:   0x3,
		name: "hello",
		desc: "lorem ipsum",
	},
}

func processBadges(a Activity, p Player) Player {
	// for i := range Badges {
	// 	if Badges[i].id
	// }

	var remainingBadges []Badge

	for _, v := range Badges {
		var foundMatch = false
		for _, vp := range p.badges {
			if vp.id == v.id {
				foundMatch = true
				break
			}
		}
		if !foundMatch {
			remainingBadges = append(remainingBadges, v)
		}
	}

	fmt.Println(remainingBadges)

	return p
}
