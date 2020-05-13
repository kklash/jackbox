package jackbox

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	// ErrEmptyRoom is returned (wrapped) by FetchRoomInfo if
	// the room is not found to exist.
	ErrEmptyRoom = errors.New("Room code is empty")

	// ErrInvalidRoomInfo is returned (wrapped) by FetchRoomInfo or ParseRoomInfo
	// if the room info is returned in an invalid format.
	ErrInvalidRoomInfo = errors.New("Room info is malformed")
)

// Response returned by jackbox servers when an active room is found.
// If no room is found, only a { success, error } object is returned.
type RoomInfo struct {
	App               string `json:"apptag"`
	AppId             string `json:"appid"`
	AudienceEnabled   bool   `json:"audienceEnabled"`
	AudienceMembers   int    `json:"numAudience"`
	JoinAs            string `json:"joinAs"`
	PasswordProtected bool   `json:"requiresPassword"`
	RoomCode          string `json:"roomid"`
	Server            string `json:"server"`
}

// Fetch the RoomInfo for a specific room code. Returns ErrEmptyRoom
// if the room code doesn't have an active room.
func FetchRoomInfo(roomCode string) (*RoomInfo, error) {
	url := API_URL("room", roomCode)
	resp, err := http.Get(url.String())
	if err != nil {
		return nil, tracerr("FetchRoomInfo: GET request", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, ErrEmptyRoom
	}

	return ParseRoomInfo(resp.Body)
}

// Parse a RoomInfo from a reader.
func ParseRoomInfo(r io.Reader) (*RoomInfo, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, tracerr("ParseRoomInfo: reading from reader", err)
	}

	roomInfo := new(RoomInfo)
	if err = json.Unmarshal(data, roomInfo); err != nil {
		return nil, tracerr("ParseRoomInfo: decoding response", err)
	}

	if !IsValidRoomInfo(roomInfo) {
		return nil, tracerr("ParseRoomInfo: checking room info validity", ErrInvalidRoomInfo)
	}

	return roomInfo, nil
}

// IsValidRoomInfo determines if the roomInfo is complete and can be
// used as a valid source of truth for information about a room.
func IsValidRoomInfo(roomInfo *RoomInfo) bool {
	return roomInfo != nil &&
		roomInfo.App != "" &&
		roomInfo.AppId != "" &&
		roomInfo.JoinAs != "" &&
		roomInfo.RoomCode != ""
}
