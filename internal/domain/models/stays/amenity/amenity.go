package amenity

type Amenity int // @name Amenity represents the type for the amenities enumeration

// Enumerate the available amenities
const (
	Wifi Amenity = iota
	AirConditioner
	PetsAllowed
	BreakfastIncluded
	VacuumCleaner
	WorkingArea
	WashingMachine
	TV
	HomeLightControl
	SmartDoorLocks
	VoiceAssistant
	TouchControlPanels
)

// String method to convert Amenity to its string representation
func (a Amenity) String() string {
	return [...]string{
		"Wi-fi",
		"Air conditioner",
		"Pets allowed",
		"Breakfast included",
		"Vacuum cleaner",
		"Working area",
		"Washing machine",
		"TV",
		"Home light control",
		"Smart door locks",
		"Voice assistant",
		"Touch control panels",
	}[a]
}
