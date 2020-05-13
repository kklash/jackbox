package jackbox

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

type ParseRoomInfoFixture struct {
	input    string
	roomInfo *RoomInfo
	valid    bool
}

func TestParseRoomInfo(t *testing.T) {
	fixtures := []*ParseRoomInfoFixture{
		{
			input:    "{}",
			roomInfo: nil,
			valid:    false,
		},
		{
			input: `{"roomid":"JFID","server":"ecast.jackboxgames.com","apptag":"auction","appid":"imanappid","numAudience":0,"audienceEnabled":false,"joinAs":"player","requiresPassword":false}`,
			roomInfo: &RoomInfo{
				App:               "auction",
				AppId:             "imanappid",
				AudienceEnabled:   false,
				AudienceMembers:   0,
				JoinAs:            "player",
				PasswordProtected: false,
				RoomCode:          "JFID",
				Server:            "ecast.jackboxgames.com",
			},
			valid: true,
		},
	}

	for _, fixture := range fixtures {
		buf := strings.NewReader(fixture.input)
		roomInfo, err := ParseRoomInfo(buf)
		if fixture.valid && err != nil {
			t.Errorf("Unexpected error parsing room info: %s", err)
			continue
		} else if !fixture.valid && err == nil {
			t.Errorf("Expected error parsing invalid room info JSON, received none")
			continue
		}

		if !reflect.DeepEqual(roomInfo, fixture.roomInfo) {
			actual, _ := json.Marshal(roomInfo)
			expected, _ := json.Marshal(fixture.roomInfo)
			t.Errorf("Room info does not match\nWanted %s\nGot    %s", string(expected), string(actual))
			continue
		}
	}
}
