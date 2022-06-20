package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/migrate"
	"testing"
)

func TestResolver_APIGetCMSWorkoutSets(t *testing.T) {
	//設定 migrate
	if err := migrate.Mock().Up(nil); err != nil {
		t.Fatalf(err.Error())
	}
	defer func() {
		if err := migrate.Mock().Down(nil); err != nil {
			t.Fatalf(err.Error())
		}
	}()
	
}
