package db
import (
	"testing"
	"git.neds.sh/matty/entain/racing/proto/racing"
)

func TestListRaces_VisibleOnly(t *testing.T) {
	// setup in-memory SQLite DB
	// seed with a mix of visible and invisible races
	// call repo.List(&racing.ListRacesRequestFilter{Visible: true})
	// assert all returned races have Visible == true

	db := NewRacesRepo(nil)
	db.Init()

	db.List(&racing.ListRacesRequestFilter{Visible: true})

	assert.Equal(t, len(races), 10)
	for _, race := range races {
		assert.True(t, race.Visible)
	}
}