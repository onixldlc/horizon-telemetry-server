package utils

var CarDashDictionary = map[string]string{
	"PositionX":                   "f",
	"PositionY":                   "f",
	"PositionZ":                   "f",
	"Speed":                       "f",
	"Power":                       "f",
	"Torque":                      "f",
	"TireTempFrontLeft":           "f",
	"TireTempFrontRight":          "f",
	"TireTempRearLeft":            "f",
	"TireTempRearRight":           "f",
	"Boost":                       "f",
	"Fuel":                        "f",
	"DistanceTraveled":            "f",
	"BestLapTime":                 "f",
	"LastLapTime":                 "f",
	"CurrentLapTime":              "f",
	"CurrentRaceTime":             "f",
	"LapNumber":                   "H", // uint16
	"RacePosition":                "B", // uint8
	"Accel":                       "B", // uint8
	"Brake":                       "B", // uint8
	"Clutch":                      "B", // uint8
	"HandBrake":                   "B", // uint8
	"Gear":                        "b", // int8
	"Steer":                       "b", // int8
	"NormalizedDrivingLine":       "b", // int8
	"NormalizedAIBrakeDifference": "b", // int8
}
