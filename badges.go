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

type Badge struct {
	id                           uint32
	name                         string
	desc                         string
	hasSingleActivityRequirement bool
	singleActivityRequirement    SingleActivityRequirement
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
		id:   0x3,
		name: "hello",
		desc: "lorem ipsum",
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

	// Iterate through badges and test them
	for _, b := range remainingBadges {
		gotBadge := testBadgeRequirements(p, a, b)
		if gotBadge {
			p.badges = append(p.badges, b)
			fmt.Println("Player was awarded badge", b.name, "[", b.desc, "]")
		}
	}

	fmt.Println(remainingBadges)

	return p
}

func testBadgeRequirements(p Player, a ActivityEntry, b Badge) bool {
	// Process remaining badgess
	req := b.singleActivityRequirement

	activityTypeRequirement := true
	if req.hasActivityTypeRequirement {
		activityTypeRequirement = a.activity == req.activityTypeRequirement
	}

	// process distance
	distOk := a.distance >= req.distance

	// process temp
	tempOk := false
	if req.hasTemperatureRequirement {
		if req.temperature.greaterThan {
			tempOk = a.outsiteTemp > req.temperature.temperature
		} else {
			tempOk = a.outsiteTemp < req.temperature.temperature
		}
	}

	return activityTypeRequirement && distOk && tempOk
}