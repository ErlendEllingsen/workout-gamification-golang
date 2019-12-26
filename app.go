package main

import "fmt"
import "io/ioutil"

type Player struct {
	level    int32
	progress float32
	badges   []Badge
	stats    PlayerStats
}

type PlayerStats struct {
	totalRunningKm float32
}

type Activity struct {
	id     string
	name   string
	cardio bool
}

type ActivitiesStruct struct {
	skiing_xc Activity
	running   Activity
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
		cardio: true,
	},
	running: Activity{
		id:     "running",
		name:   "Running",
		cardio: true,
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
	// append stats
	p.stats.totalRunningKm += a.distance
	// calculate progress for rank
	progressFromActivity := a.distance * 1.2
	return p.AddRank(progressFromActivity)
}

func main() {

	localMode := true
	var acts []StravaActivity

	if localMode {
		jsonRaw, _ := ioutil.ReadFile("strava-data.json")
		acts = parseActivities(jsonRaw)
	} else {
		acts = getActivitiesFromStrava(false)
	}

	// fmt.Println(acts)

	var p = Player{
		level:  1,
		badges: []Badge{},
	}

	// Add activities to player
	for _, sa := range acts {
		a := sa.ToActivityEntry()
		p = p.AddActivity(a)
		p = p.ProcessBadges(a)
	}

	// fmt.Println(p)

	// act := ActivityEntry{
	// 	activity:    Activities.skiing_xc,
	// 	distance:    16.0,
	// 	outsiteTemp: -1.0,
	// }

	// p = p.AddActivity(act)
	// p = p.ProcessBadges(act)

	// fmt.Println(p)

}
