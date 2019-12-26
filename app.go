package main

import "encoding/json"
import "fmt"

// import "os"
import "io/ioutil"

type Player struct {
	Level    int32
	Progress float32
	Badges   []uint32
	Stats    PlayerStats
}

type PlayerStats struct {
	TotalRunningKm float32
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
	targetProgress := float32(p.Level) * 2.5
	newProgress := p.Progress + newProg

	if targetProgress <= newProgress {
		p.Level++
		p.Progress = 0
		surplusProgress := newProgress - targetProgress
		fmt.Println("Player level up", p.Level, "surplus progress", surplusProgress)
		if surplusProgress > 0 {
			return p.AddRank(surplusProgress)
		}
	} else {
		p.Progress = newProgress
		fmt.Println("Player level", p.Level, "current progress", p.Progress)
	}

	return p
}

func (p Player) AddActivity(a ActivityEntry) Player {
	// append stats
	p.Stats.TotalRunningKm += a.distance
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
		Level:  1,
		Badges: []uint32{},
	}

	// Add activities to player
	for _, sa := range acts {
		fmt.Println("Activity", sa.ID, sa.Distance)
		a := sa.ToActivityEntry()
		p = p.AddActivity(a)
		p = p.ProcessBadges(a)
	}

	dat, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(dat))

	// fmt.Println(p)

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
