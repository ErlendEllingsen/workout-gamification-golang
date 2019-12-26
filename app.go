package main

import "fmt"

type Player struct {
	level    int
	progress float32
	badges   []Badge
}

type Activity struct {
	cardio   bool
	distance float32
}

type ActivitiesStruct struct {
	skiing_xc Activity
}

var Activities = ActivitiesStruct{
	skiing_xc: Activity{
		cardio:   false,
		distance: 5.0,
	},
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
	var p = Player{
		level: 1,
		badges: []Badge{
			Badge{
				id:   0x2,
				name: "Skier",
				desc: "Went on a ski trip",
			},
		},
	}

	fmt.Println(p)

	p = p.AddActivity(Activities.skiing_xc)

	fmt.Println(p)

	processBadges(Activities.skiing_xc, p)

}
