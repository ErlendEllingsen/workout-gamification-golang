package main

import "fmt"

type TemperatureRequirement struct {
	temperature float32
	greaterThan bool
}

type SingleActivityRequirement struct {
	distance                   float32
	hasTemperatureRequirement  bool
	temperature                TemperatureRequirement
	hasActivityTypeRequirement bool
	activityTypeRequirement    Activity
}

type StatsRequirement struct {
	hasDistanceRanRequirement bool
	distanceRanRequirement    float32
}

type PlayerRequirement struct {
	hasLevelRequirement bool
	levelRequirement    int32
	hasStatsRequirement bool
	statsRequirement    StatsRequirement
}

type Badge struct {
	id                           uint32
	name                         string
	desc                         string
	hasSingleActivityRequirement bool
	singleActivityRequirement    SingleActivityRequirement
	hasPlayerRequirement         bool
	playerRequirement            PlayerRequirement
}

var Badges = []Badge{
	Badge{
		id:                           0x1,
		name:                         "Hello winter",
		desc:                         "Complete an activity in freezing temperatures",
		hasSingleActivityRequirement: true,
		singleActivityRequirement: SingleActivityRequirement{
			distance:                  0.0,
			hasTemperatureRequirement: true,
			temperature: TemperatureRequirement{
				greaterThan: false,
				temperature: 0.0,
			},
		},
	},
	Badge{
		id:                           0x2,
		name:                         "Skier basic boomer",
		desc:                         "Ski for at least 5km",
		hasSingleActivityRequirement: true,
		singleActivityRequirement: SingleActivityRequirement{
			hasActivityTypeRequirement: true,
			activityTypeRequirement:    Activities.skiing_xc,
			distance:                   5.0,
			hasTemperatureRequirement:  false,
			temperature:                TemperatureRequirement{},
		},
	},
	Badge{
		id:                   0x3,
		name:                 "Gettin' started",
		desc:                 "Reached level 3",
		hasPlayerRequirement: true,
		playerRequirement: PlayerRequirement{
			hasLevelRequirement: true,
			levelRequirement:    3,
		},
	},
	Badge{
		id:                   0x4,
		name:                 "Rising star",
		desc:                 "Reached level 5",
		hasPlayerRequirement: true,
		playerRequirement: PlayerRequirement{
			hasLevelRequirement: true,
			levelRequirement:    5,
		},
	},
	Badge{
		id:                   0x5,
		name:                 "RiSiNg bOomer",
		desc:                 "Reached level 10",
		hasPlayerRequirement: true,
		playerRequirement: PlayerRequirement{
			hasLevelRequirement: true,
			levelRequirement:    10,
		},
	},
	Badge{
		id:                           0x6,
		name:                         "Baby steps",
		desc:                         "Ran 1km",
		hasSingleActivityRequirement: true,
		singleActivityRequirement: SingleActivityRequirement{
			distance:                   1.0,
			hasActivityTypeRequirement: true,
			activityTypeRequirement:    Activities.running,
		},
	},
	Badge{
		id:                           0x7,
		name:                         "Sweaty shoes",
		desc:                         "Ran 5km - we're getting somewhere!",
		hasSingleActivityRequirement: true,
		singleActivityRequirement: SingleActivityRequirement{
			distance:                   5.0,
			hasActivityTypeRequirement: true,
			activityTypeRequirement:    Activities.running,
		},
	},
	Badge{
		id:                           0x8,
		name:                         "Rocket fueled sweaty shoes",
		desc:                         "Ran 10km - thats insane man!",
		hasSingleActivityRequirement: true,
		singleActivityRequirement: SingleActivityRequirement{
			distance:                   10.0,
			hasActivityTypeRequirement: true,
			activityTypeRequirement:    Activities.running,
		},
	},
	Badge{
		id:                   0x9,
		name:                 "Around de earth",
		desc:                 "Ran 30km total",
		hasPlayerRequirement: true,
		playerRequirement: PlayerRequirement{
			hasStatsRequirement: true,
			statsRequirement: StatsRequirement{
				hasDistanceRanRequirement: true,
				distanceRanRequirement:    30.0,
			},
		},
	},
	Badge{
		id:                   0x10,
		name:                 "Around de moon",
		desc:                 "Ran 50km total",
		hasPlayerRequirement: true,
		playerRequirement: PlayerRequirement{
			hasStatsRequirement: true,
			statsRequirement: StatsRequirement{
				hasDistanceRanRequirement: true,
				distanceRanRequirement:    50.0,
			},
		},
	},
}

func (p Player) ProcessBadges(a ActivityEntry) Player {
	// for i := range Badges {
	// 	if Badges[i].id
	// }

	var remainingBadges []Badge

	// Identify remaining badges
	for _, v := range Badges {
		var foundMatch = false
		for _, vp := range p.Badges {
			if vp == v.id {
				foundMatch = true
				break
			}
		}
		if !foundMatch {
			remainingBadges = append(remainingBadges, v)
		}
	}

	// Iterate through badges and test them
	for _, b := range remainingBadges {
		gotBadge := testBadgeRequirements(p, a, b)
		if gotBadge {
			p.Badges = append(p.Badges, b.id)
			fmt.Println("Player was awarded badge", b.name, "[", b.desc, "]")
		}
	}

	// fmt.Println(remainingBadges)

	return p
}

func testBadgeRequirements(p Player, a ActivityEntry, b Badge) bool {
	// Process remaining badgess
	req := b.singleActivityRequirement

	activityTypeRequirement := true
	if req.hasActivityTypeRequirement {
		activityTypeRequirement = a.activity.id == req.activityTypeRequirement.id
	}

	// process distance
	distOk := a.distance >= req.distance

	// process temp
	tempOk := true
	if req.hasTemperatureRequirement {
		if req.temperature.greaterThan {
			tempOk = a.outsiteTemp > req.temperature.temperature
		} else {
			tempOk = a.outsiteTemp < req.temperature.temperature
		}
	}

	// combine activity reqs
	actionRequirements := activityTypeRequirement && distOk && tempOk

	// process player requirements
	playerRequirements := true
	if b.hasPlayerRequirement {
		playerReq := b.playerRequirement
		levelRequirement := true
		if playerReq.hasLevelRequirement {
			levelRequirement = p.Level >= playerReq.levelRequirement
		}

		statsRequirement := true
		if playerReq.hasStatsRequirement {
			statsReq := playerReq.statsRequirement
			distanceRanReq := true
			if statsReq.hasDistanceRanRequirement {
				distanceRanReq = p.Stats.TotalRunningKm >= statsReq.distanceRanRequirement
			}
			statsRequirement = distanceRanReq
		}

		playerRequirements = levelRequirement && statsRequirement
	}

	return actionRequirements && playerRequirements
}
