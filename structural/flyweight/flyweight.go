/*
The Flyweight design pattern allows sharing state of a heavy object between many instances
of some type. We can share all possible states of objects in a single common object, and
thus minimize object creation by using pointers to already created objects.
*/

package flyweight

type FlyweightFactory interface {
	GetTeam(uint64) *Team
	GetNumberOfObjects() int
}

type teamFlyweightFactory struct {
	createdTeams map[uint64]*Team
}

func NewTeamFactory() FlyweightFactory {
	return &teamFlyweightFactory{
		createdTeams: make(map[uint64]*Team, 0),
	}
}

func (f *teamFlyweightFactory) GetTeam(teamID uint64) *Team {
	if team, ok := f.createdTeams[teamID]; ok {
		return team
	}

	team := getTeamFactory(teamID)
	f.createdTeams[teamID] = &team
	return &team
}

func (f *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(f.createdTeams)
}

func getTeamFactory(teamID uint64) Team {
	switch teamID {
	case TEAM_B:
		return Team{
			ID:   2,
			Name: "TeamB",
		}
	default:
		return Team{
			ID:   1,
			Name: "TeamA",
		}
	}
}
