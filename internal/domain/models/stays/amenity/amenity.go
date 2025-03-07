package amenity

// Amenity represents the type for the amenities enumeration
type Amenity string

// Enumerate the available amenities
const (
	Wifi               Amenity = "Wi-fi"
	AirConditioner     Amenity = "Air conditioner"
	PetsAllowed        Amenity = "Pets allowed"
	BreakfastIncluded  Amenity = "Breakfast included"
	VacuumCleaner      Amenity = "Vacuum cleaner"
	WorkingArea        Amenity = "Working area"
	WashingMachine     Amenity = "Washing machine"
	TV                 Amenity = "TV"
	HomeLightControl   Amenity = "Home light control"
	SmartDoorLocks     Amenity = "Smart door locks"
	VoiceAssistant     Amenity = "Voice assistant"
	TouchControlPanels Amenity = "Touch control panels"
)

// String method to convert Amenity to its string representation
func (a Amenity) String() string {
	return string(a)
}

// AllAmenities returns a slice of all available amenities
func AllAmenities() []Amenity {
	return []Amenity{
		Wifi,
		AirConditioner,
		PetsAllowed,
		BreakfastIncluded,
		VacuumCleaner,
		WorkingArea,
		WashingMachine,
		TV,
		HomeLightControl,
		SmartDoorLocks,
		VoiceAssistant,
		TouchControlPanels,
	}
}
