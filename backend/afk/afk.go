package afk

import (
	"fmt"

	"mugowalker/backend/afk/activities"
	"mugowalker/backend/afk/repository"
	"mugowalker/backend/pilot"
)

type Game struct {
	Name      string
	Active    bool
	Locations []any
	User      *repository.User
	profile   *pilot.Pilot
}

func (g *Game) String() string {
	return fmt.Sprintf("Name: %v\n User:%v\n", g.Name, g.User.Username)
}

// New Game for a given User
func New(up *pilot.Pilot) *Game {

	anylocs := activities.AllLocations()
	for _, l := range anylocs {
		if loc, ok := l.(activities.Location); ok {
			for _, kw := range loc.Keywords() {
				if kw == "%account" {
					loc.Kws = append(loc.Kws, up.Account)
				}
			}

		}
		anylocs = append(anylocs, l)
	}
	user := &repository.User{Username: up.Account} //repository.GetUser(up.Account)

	return &Game{
		Name:      up.Game,
		Locations: anylocs,
		Active:    true,
		User:      user,
		profile:   up,
	}
}

/*
	|			|
pt. |Quest  	| %b
-----------------------
10  |Loot x2	|	1
10	|FastReward	|	1
10  |Friendship	|	1
10	|Wrizz		|	1
20	|Arena1x1	|	1
10  |Inn		|	1
20	|Fight Camp	|	0
10	|Fight KT	|	1


hard to implement
10 	|Bounty		|
20	|summon		|
	|ArenaTopEnemy
	|FRqty		|
// */
