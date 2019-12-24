package main

import "fmt"

type Badge struct {
	name string
	desc string
}

type Player struct {
	level    int
	progress float32
	badges   []Badge
}

type Activity struct {
	cardio   bool
	distance float32
}

func (p Player) AddRank(newProg float32) Player {
	targetProgress := float32(p.level) * 2.5
	newProgress := p.progress + newProg

	if targetProgress <= newProgress {
		p.level++
		p.progress = 0
		surplusProgress := newProgress - targetProgress
		if surplusProgress > 0 {
			return p.AddRank(surplusProgress)
		}
	}

	return p
}

func (p Player) AddActivity(a Activity) Player {
	progressFromActivity := a.distance * 1.2
	return p.AddRank(progressFromActivity)
}

func main() {
	var x = Player{
		level: 1,
		badges: []Badge{
			Badge{
				name: "Skier",
				desc: "Went on a ski trip",
			},
		},
	}

	fmt.Println(x)

	x = x.AddActivity(Activity{
		cardio:   false,
		distance: 50.0,
	})

	fmt.Println(x)

}
