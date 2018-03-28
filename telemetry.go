package gopubg

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"time"
)

// TelemetryEvent represents any event from a telemetry file
type TelemetryEvent struct {
	// Common fields
	Version   int       `json:"_V"`
	Timestamp time.Time `json:"_D"`
	Type      string    `json:"_T"`
	U         bool      `json:"_U"`

	// --- Player
	// Events: LogPlayerLogin, LogPlayerLogout, LogPlayerCreate, LogPlayerPosition, LogPlayerAttack, LogPlayerTakeDamage, LogPlayerKill
	Result             bool                `json:"result"`
	ErrorMessage       string              `json:"errorMessage"`
	AccountID          string              `json:"accountId"`
	Character          *TelemetryCharacter `json:"character"`
	ElapsedTime        float64             `json:"elapsedTime"`
	NumAlivePlayers    int                 `json:"numAlivePlayers"`
	AttackID           int                 `json:"attackId"`
	Attacker           *TelemetryCharacter `json:"attacker"`
	AttackType         string              `json:"attackType"`
	Weapon             *TelemetryItem      `json:"weapon"`
	Vehicle            *TelemetryVehicle   `json:"vehicle"`
	Victim             *TelemetryCharacter `json:"victim"`
	DamageTypeCategory string              `json:"damageTypeCategory"`
	DamageReason       string              `json:"damageReason"`
	Damage             float64             `json:"damage"`
	DamageCauserName   string              `json:"damageCauserName"`
	Killer             *TelemetryCharacter `json:"killer"`
	Distance           float64             `json:"distance"`

	// --- Vehicle
	// Events: LogVehicleRide, LogVehicleLeave
	// Character already defined
	// Vehicle already defined

	// --- Item
	// Events: LogItemPickup, LogItemEquip, LogItemUnequip, LogItemAttach, LogItemDrop, LogItemDetach, LogItemUse
	Item       *TelemetryItem `json:"item"`
	ParentItem *TelemetryItem `json:"parentItem"`
	ChildItem  *TelemetryItem `json:"childItem"`

	// --- Match
	// Events: LogMatchStart, LogMatchEnd, LogMatchDefinition
	Characters  []*TelemetryCharacter
	MatchID     string `json:"matchId"`
	PingQuality string `json:"pingQuality"`

	// --- Care package
	// Events: LogCarePackageSpawn, LogCarePackageLand
	ItemPackage *TelemetryItemPackage `json:"itemPackage"`

	// --- Game
	// Events: LogGameStatePeriodic
	GameState *TelemetryGameState
}

// TelemetryItemPackage represents an item package
type TelemetryItemPackage struct {
	ItemPackageID string             `json:"itemPackageId"`
	Location      *TelemetryLocation `json:"location"`
	Items         []*TelemetryItem   `json:"items"`
}

// TelemetryGameState represents the state of a game
type TelemetryGameState struct {
	ElapsedTime              int                `json:"elapsedTime"`
	NumAliveTeams            int                `json:"numAliveTeams"`
	NumJoinPlayers           int                `json:"numJoinPlayers"`
	NumStartPlayers          int                `json:"numStartPlayers"`
	NumAlivePlayers          int                `json:"numAlivePlayers"`
	SafetyZonePosition       *TelemetryLocation `json:"safetyZonePosition"`
	SafetyZoneRadius         float64            `json:"safetyZoneRadius"`
	PoisonGasWarningPosition *TelemetryLocation `json:"poisonGasWarningPosition"`
	PoisonGasWarningRadius   float64            `json:"poisonGasWarningRadius"`
	RedZonePosition          *TelemetryLocation `json:"redZonePosition"`
	RedZoneRadius            float64            `json:"redZoneRadius"`
}

// TelemetryVehicle represents a vehicle
type TelemetryVehicle struct {
	VehicleType   string  `json:"vehicleType"`
	VehicleID     string  `json:"vehicleId"`
	HealthPercent float64 `json:"healthPercent"`
	FuelPercent   float64 `json:"feulPercent"`
}

// TelemetryItem represents an item
type TelemetryItem struct {
	ItemID        string   `json:"itemId"`
	StackCount    int      `json:"stackCount"`
	Category      string   `json:"category"`
	SubCategory   string   `json:"subCategory"`
	AttachedItems []string `json:"attachedItems"`
}

// TelemetryCharacter represents a character
type TelemetryCharacter struct {
	Name      string             `json:"name"`
	TeamID    int                `json:"teamId"`
	Health    float64            `json:"health"`
	Location  *TelemetryLocation `json:"location"`
	Ranking   int                `json:"ranking"`
	AccountID string             `json:"accountId"`
}

// TelemetryLocation represents a location
type TelemetryLocation struct {
	X float64 `json:"X"`
	Y float64 `json:"Y"`
	Z float64 `json:"Z"`
}

// Telemetry represents the context of a telemetry file
type Telemetry struct {
	Events      []*TelemetryEvent
	PlayerNames []string
}

func newTelemetry() *Telemetry {
	return &Telemetry{
		Events:      make([]*TelemetryEvent, 0),
		PlayerNames: make([]string, 0),
	}
}

// ParseTelemetry parses a json response containing telemetry information
func ParseTelemetry(in io.Reader) (*Telemetry, error) {
	t := newTelemetry()

	data, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}

	// Parse events
	if err := json.Unmarshal(data, &t.Events); err != nil {
		return nil, err
	}

	// Find players
	for _, e := range t.Events {
		if e.Type == "LogMatchStart" {
			for _, c := range e.Characters {
				t.PlayerNames = append(t.PlayerNames, c.Name)
			}
			break
		}
	}

	return t, nil
}
