package gopubg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"time"

	"github.com/sirupsen/logrus"
)

// TelemetryEventType represents the type of a telemetry event
type TelemetryEventType int

func findIndex(key string, possibleKeys []string) int {
	for idx, possibleKey := range possibleKeys {
		if key == possibleKey {
			return idx
		}
	}
	return -1
}

// Telemetry event types
const (
	PlayerLogin TelemetryEventType = iota
	PlayerLogout
	PlayerCreate
	PlayerPosition
	PlayerAttack
	PlayerTakeDamage
	PlayerKill
	ItemPickup
	ItemDrop
	ItemEquip
	ItemUnequip
	ItemAttach
	ItemDetach
	ItemUse
	VehicleRide
	VehicleLeave
	VehicleDestroy
	MatchStart
	MatchEnd
	MatchDefinition
	GameStatePeriodic
	CarePackageSpawn
	CarePackageLand
)

// KnownEventTypes represents supported types
var KnownEventTypes = []string{
	"LogPlayerLogin",
	"LogPlayerLogout",
	"LogPlayerCreate",
	"LogPlayerPosition",
	"LogPlayerAttack",
	"LogPlayerTakeDamage",
	"LogPlayerKill",
	"LogItemPickup",
	"LogItemDrop",
	"LogItemEquip",
	"LogItemUnequip",
	"LogItemAttach",
	"LogItemDetach",
	"LogItemUse",
	"LogVehicleRide",
	"LogVehicleLeave",
	"LogVehicleDestroy",
	"LogMatchStart",
	"LogMatchEnd",
	"LogMatchDefinition",
	"LogGameStatePeriodic",
	"LogCarePackageSpawn",
	"LogCarePackageLand",
}

// UnmarshalJSON verifies the type of a telemetry event is known
func (t *TelemetryEventType) UnmarshalJSON(data []byte) error {
	if t == nil {
		return errors.New("TelemetryEventType: UnmarshalJSON on nil pointer")
	}

	key := string(data)
	idx := findIndex(key[1:len(key)-1], KnownEventTypes)
	if idx == -1 {
		return fmt.Errorf("TelemetryEventType: Unknown type %s", key)
	}

	*t = TelemetryEventType(idx)
	return nil
}

// TelemetryAttackType represents the type of an attack
type TelemetryAttackType int

// Telemetry attack types
const (
	AttackTypeRedZone TelemetryAttackType = iota
	AttackTypeWeapon
)

var knownAttackTypes = []string{
	"RedZone",
	"Weapon",
}

// UnmarshalJSON verifies the type of an attack is known
func (t *TelemetryAttackType) UnmarshalJSON(data []byte) error {
	if t == nil {
		return errors.New("TelemetryAttackType: UnmarshalJSON on nil pointer")
	}

	key := string(data)
	idx := findIndex(key[1:len(key)-1], knownAttackTypes)
	if idx == -1 {
		return fmt.Errorf("TelemetryAttackType: Unknown type %s", key)
	}

	*t = TelemetryAttackType(idx)
	return nil
}

// TelemetrySubCategory represents the category of an item
type TelemetrySubCategory int

// Telemetry sub categories
const (
	SubCategoryBackpack TelemetrySubCategory = iota
	SubCategoryBoost
	SubCategoryFuel
	SubCategoryHandgun
	SubCategoryHeadgear
	SubCategoryHeal
	SubCategoryMain
	SubCategoryMelee
	SubCategoryThrowable
	SubCategoryVest
	SubCategoryJacket
	SubCategoryNone
	SubCategoryEmpty
)

var knownSubCategories = []string{
	"Backpack",
	"Boost",
	"Fuel",
	"Handgun",
	"Headgear",
	"Heal",
	"Main",
	"Melee",
	"Throwable",
	"Vest",
	"Jacket",
	"None",
	"",
}

// UnmarshalJSON verifies the type of a subcategory
func (t *TelemetrySubCategory) UnmarshalJSON(data []byte) error {
	if t == nil {
		return errors.New("TelemetrySubCategory: UnmarshalJSON on nil pointer")
	}

	key := string(data)
	idx := findIndex(key[1:len(key)-1], knownSubCategories)
	if idx == -1 {
		return fmt.Errorf("TelemetrySubCategory: Unknown type %s", key)
	}

	*t = TelemetrySubCategory(idx)
	return nil
}

// TelemetryDamageType represents the different types of damage
type TelemetryDamageType int

// Telemetry damage types
const (
	DamageBlueZone TelemetryDamageType = iota
	DamageDrown
	DamageExplosionGrenade
	DamageExplosionRedZone
	DamageExplosionVehicle
	DamageGroggy
	DamageGun
	DamageInstantFall
	DamageMelee
	DamageMolotov
	DamageVehicleCrashHit
	DamageVehicleHit
	DamageEmpty
)

var knownDamageTypes = []string{
	"Damage_BlueZone",
	"Damage_Drown",
	"Damage_Explosion_Grenade",
	"Damage_Explosion_RedZone",
	"Damage_Explosion_Vehicle",
	"Damage_Groggy",
	"Damage_Gun",
	"Damage_Instant_Fall",
	"Damage_Melee",
	"Damage_Molotov",
	"Damage_VehicleCrashHit",
	"Damage_VehicleHit",
	"",
}

// UnmarshalJSON verifies the type of a subcategory
func (t *TelemetryDamageType) UnmarshalJSON(data []byte) error {
	if t == nil {
		return errors.New("TelemetryDamageType: UnmarshalJSON on nil pointer")
	}

	key := string(data)
	idx := findIndex(key[1:len(key)-1], knownDamageTypes)
	if idx == -1 {
		return fmt.Errorf("TelemetryDamageType: Unknown type %s", key)
	}

	*t = TelemetryDamageType(idx)
	return nil
}

// TelemetryDamageReason represents the reason of the damage
type TelemetryDamageReason int

// Telemetry damage reasons
const (
	DamageReasonArmShot TelemetryDamageReason = iota
	DamageReasonHeadShot
	DamageReasonLegShot
	DamageReasonPelvisShot
	DamageReasonTorsoShot
	DamageReasonNonSpecific
	DamageReasonNone
)

var knownDamageReasons = []string{
	"ArmShot",
	"HeadShot",
	"LegShot",
	"PelvisShot",
	"TorsoShot",
	"NonSpecific",
	"None",
}

// UnmarshalJSON verifies the type of a subcategory
func (t *TelemetryDamageReason) UnmarshalJSON(data []byte) error {
	if t == nil {
		return errors.New("TelemetryDamageReason: UnmarshalJSON on nil pointer")
	}

	key := string(data)
	idx := findIndex(key[1:len(key)-1], knownDamageReasons)
	if idx == -1 {
		return fmt.Errorf("TelemetryDamageReason: Unknown type %s", key)
	}

	*t = TelemetryDamageReason(idx)
	return nil
}

// TelemetryEvent represents any event from a telemetry file
type TelemetryEvent struct {
	// Common fields
	Version   int                `json:"_V"`
	Timestamp time.Time          `json:"_D"`
	Type      TelemetryEventType `json:"_T"`
	U         bool               `json:"_U"`

	// --- Player
	// Events: LogPlayerLogin, LogPlayerLogout, LogPlayerCreate, LogPlayerPosition, LogPlayerAttack, LogPlayerTakeDamage, LogPlayerKill
	Result             bool                  `json:"result"`
	ErrorMessage       string                `json:"errorMessage"`
	AccountID          string                `json:"accountId"`
	Character          *TelemetryCharacter   `json:"character"`
	ElapsedTime        float64               `json:"elapsedTime"`
	NumAlivePlayers    int                   `json:"numAlivePlayers"`
	AttackID           int                   `json:"attackId"`
	Attacker           *TelemetryCharacter   `json:"attacker"`
	AttackType         TelemetryAttackType   `json:"attackType"`
	Weapon             *TelemetryItem        `json:"weapon"`
	Vehicle            *TelemetryVehicle     `json:"vehicle"`
	Victim             *TelemetryCharacter   `json:"victim"`
	DamageTypeCategory TelemetryDamageType   `json:"damageTypeCategory"`
	DamageReason       TelemetryDamageReason `json:"damageReason"`
	Damage             float64               `json:"damage"`
	DamageCauserName   string                `json:"damageCauserName"`
	Killer             *TelemetryCharacter   `json:"killer"`
	Distance           float64               `json:"distance"`

	// --- Vehicle
	// Events: LogVehicleRide, LogVehicleLeave, VehicleDestroy
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
	ItemID        string               `json:"itemId"`
	StackCount    int                  `json:"stackCount"`
	Category      string               `json:"category"`
	SubCategory   TelemetrySubCategory `json:"subCategory"`
	AttachedItems []string             `json:"attachedItems"`
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

// Player represents a player
type Player struct {
	Name      string
	AccountID string
	TeamID    int
	Events    []*TelemetryEvent
	Locations []*TelemetryLocation
	Ranking   int
}

func newPlayer(name, accountID string) *Player {
	return &Player{
		Name:      name,
		AccountID: accountID,
		TeamID:    -1,
		Events:    make([]*TelemetryEvent, 0),
		Locations: make([]*TelemetryLocation, 0),
		Ranking:   -1,
	}
}

// Telemetry represents the context of a telemetry file
type Telemetry struct {
	Events       []*TelemetryEvent
	Players      map[string]*Player
	MatchStarted bool
	PingQuality  string
	MatchID      string
}

func newTelemetry() *Telemetry {
	return &Telemetry{
		Events:       make([]*TelemetryEvent, 0),
		Players:      make(map[string]*Player),
		MatchStarted: false,
		PingQuality:  "",
		MatchID:      "",
	}
}

func (t *Telemetry) getPlayer(name, accountID string) *Player {
	if _, ok := t.Players[accountID]; !ok {
		t.Players[accountID] = newPlayer(name, accountID)
	}

	return t.Players[accountID]
}

func (t *Telemetry) addPlayerEvent(te *TelemetryEvent, character *TelemetryCharacter, matchStarted bool) {
	if character.Name == "" {
		return
	}

	player := t.getPlayer(character.Name, character.AccountID)

	if matchStarted {
		player.Events = append(player.Events, te)
		player.Locations = append(player.Locations, character.Location)
	}
}

func (t *Telemetry) processEvent(te *TelemetryEvent) {
	logrus.WithFields(logrus.Fields{
		"type": KnownEventTypes[te.Type],
	}).Debug("Processing event")

	// Look for common fields
	if te.Character != nil {
		t.addPlayerEvent(te, te.Character, t.MatchStarted)
	}

	// Look for a custom function to specialize the data
	functionName := "Process" + KnownEventTypes[te.Type]
	f := reflect.ValueOf(t).MethodByName(functionName)
	if f.IsValid() {
		f.Call([]reflect.Value{
			reflect.ValueOf(te),
		})
	}
}

// ProcessLogMatchDefinition deals with event of type MatchDefinition
func (t *Telemetry) ProcessLogMatchDefinition(te *TelemetryEvent) {
	t.PingQuality = te.PingQuality
	t.MatchID = te.MatchID
}

// ProcessLogMatchStart deals with event of type MatchStart
func (t *Telemetry) ProcessLogMatchStart(te *TelemetryEvent) {
	t.MatchStarted = true
}

// ProcessLogMatchEnd deals with event of type MatchEnd
func (t *Telemetry) ProcessLogMatchEnd(te *TelemetryEvent) {
	// Update player ranking
	for _, c := range te.Characters {
		player := t.getPlayer(c.Name, c.AccountID)
		player.Ranking = c.Ranking
	}
}

// ProcessLogPlayerCreate deals with event of type PlayerCreate
func (t *Telemetry) ProcessLogPlayerCreate(te *TelemetryEvent) {
	player := t.getPlayer(te.Character.Name, te.Character.AccountID)
	player.TeamID = te.Character.TeamID
}

// ProcessLogPlayerAttack deals with event of type PlayerAttack
func (t *Telemetry) ProcessLogPlayerAttack(te *TelemetryEvent) {
	t.addPlayerEvent(te, te.Attacker, t.MatchStarted)
}

// ProcessLogPlayerTakeDamage deals with event of type PlayerTakeDamage
func (t *Telemetry) ProcessLogPlayerTakeDamage(te *TelemetryEvent) {
	t.addPlayerEvent(te, te.Attacker, t.MatchStarted)
	t.addPlayerEvent(te, te.Victim, t.MatchStarted)
}

// ProcessLogVehicleDestroy deals with event of type VehicleDestroy
func (t *Telemetry) ProcessLogVehicleDestroy(te *TelemetryEvent) {
	t.addPlayerEvent(te, te.Attacker, t.MatchStarted)
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
		t.processEvent(e)
	}

	return t, nil
}
