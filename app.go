package main

import "fmt"

type Player struct {
	level    int
	progress float32
	badges   []Badge
}

type Activity struct {
	id     string
	name   string
	cardio bool
}

type ActivitiesStruct struct {
	skiing_xc Activity
}

type ActivityEntry struct {
	activity    Activity
	outsiteTemp float32
	distance    float32
}

var Activities = ActivitiesStruct{
	skiing_xc: Activity{
		id:     "skiing_xc",
		name:   "Cross Country",
		cardio: false,
	},
}

func (p Player) AddRank(newProg float32) Player {
	targetProgress := float32(p.level) * 2.5
	newProgress := p.progress + newProg

	if targetProgress <= newProgress {
		p.level++
		p.progress = 0
		surplusProgress := newProgress - targetProgress
		fmt.Println("Player level up", p.level, "surplus progress", surplusProgress)
		if surplusProgress > 0 {
			return p.AddRank(surplusProgress)
		}
	} else {
		p.progress = newProgress
		fmt.Println("Player level", p.level, "current progress", p.progress)
	}

	return p
}

func (p Player) AddActivity(a ActivityEntry) Player {
	progressFromActivity := a.distance * 1.2
	return p.AddRank(progressFromActivity)
}

func main() {
	var p = Player{
		level:  1,
		badges: []Badge{},
	}

	// fmt.Println(p)

	act := ActivityEntry{
		activity:    Activities.skiing_xc,
		distance:    16.0,
		outsiteTemp: -1.0,
	}

	p = p.AddActivity(act)
	p = p.ProcessBadges(act)

	// fmt.Println(p)

}
