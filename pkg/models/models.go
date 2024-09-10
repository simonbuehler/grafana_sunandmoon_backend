package models

// SunAndMoonQuery repräsentiert die Abfrageparameter für Metriken und Annotationen.
type SunAndMoonQuery struct {
	Latitude  *float64  `json:"latitude"`  // Latitude als optionaler Wert
	Longitude *float64  `json:"longitude"` // Longitude als optionaler Wert
	Target    *[]string `json:"target"`    // Zielmetriken oder Annotationen
}

// MetricDefinition definiert eine Metrik mit Titel, Text und Konfiguration.
type MetricDefinition struct {
	Title  string
	Text   string
	Config MetricConfig
}

// MetricConfig definiert die Konfiguration einer Metrik (z.B. Einheit, Dezimalstellen).
type MetricConfig struct {
	Unit     string
	Min      float64
	Decimals int
}

// AnnotationDefinition definiert eine Annotation mit Titel, Text und Tags.
type AnnotationDefinition struct {
	Title string
	Text  string
	Tags  []string
}

// SunAndMoonMetrics ist eine Map, die alle Metriken definiert.
var SunAndMoonMetrics = map[string]MetricDefinition{
	"moon_illumination": {
		Title: "Moon illumination",
		Text:  "Percentage of the moon illuminated by the sun (0.0 - 1.0)",
		Config: MetricConfig{
			Unit:     "percentunit",
			Decimals: 1,
		},
	},
	"moon_altitude": {
		Title: "Moon altitude",
		Text:  "Height of the moon in degrees (-90 - 90)",
		Config: MetricConfig{
			Unit:     "degree",
			Min:      0,
			Decimals: 1,
		},
	},
	"moon_azimuth": {
		Title: "Moon azimuth",
		Text:  "Direction of the moon along the horizon in degrees (0 - 360)",
		Config: MetricConfig{
			Unit:     "degree",
			Decimals: 1,
		},
	},
	"moon_distance": {
		Title: "Moon distance",
		Text:  "Distance to the moon in kilometers",
		Config: MetricConfig{
			Unit:     "lengthkm",
			Decimals: 0,
		},
	},
	"sun_altitude": {
		Title: "Sun altitude",
		Text:  "Height of the sun in degrees (-90 - 90)",
		Config: MetricConfig{
			Unit:     "degree",
			Min:      0,
			Decimals: 1,
		},
	},
	"sun_azimuth": {
		Title: "Sun azimuth",
		Text:  "Direction of the sun along the horizon in degrees (0 - 360)",
		Config: MetricConfig{
			Unit:     "degree",
			Decimals: 1,
		},
	},
	"sun_maximum_altitude": {
		Title: "Maximum sun altitude of the day",
		Text:  "Maximum height of the sun of the day (at solar noon) in degrees (-90 - 90)",
		Config: MetricConfig{
			Unit:     "degree",
			Min:      0,
			Decimals: 1,
		},
	},
}

// SunAndMoonAnnotations ist eine Map, die alle Annotationen definiert.
var SunAndMoonAnnotations = map[string]AnnotationDefinition{
	"sunrise": {
		Title: "Sunrise",
		Text:  "Top edge of the sun appears on the horizon",
		Tags:  []string{"sun"},
	},
	"sunriseEnd": {
		Title: "Sunrise ends",
		Text:  "Bottom edge of the sun touches the horizon",
		Tags:  []string{"sun"},
	},
	"goldenHourEnd": {
		Title: "Morning golden hour ends",
		Text:  "Soft light, best time for photography",
		Tags:  []string{"sun"},
	},
	"solarNoon": {
		Title: "Solar noon",
		Text:  "Sun is in the highest position",
		Tags:  []string{"sun"},
	},
	"goldenHour": {
		Title: "Evening golden hour starts",
		Text:  "Soft light, best time for photography",
		Tags:  []string{"sun"},
	},
	"sunsetStart": {
		Title: "Sunset starts",
		Text:  "Bottom edge of the sun touches the horizon",
		Tags:  []string{"sun"},
	},
	"sunset": {
		Title: "Sunset",
		Text:  "Sun disappears below the horizon, evening civil twilight starts",
		Tags:  []string{"sun"},
	},
	"dusk": {
		Title: "Dusk",
		Text:  "Evening nautical twilight starts",
		Tags:  []string{"sun"},
	},
	"nauticalDusk": {
		Title: "Nautical dusk",
		Text:  "Evening astronomical twilight starts",
		Tags:  []string{"sun"},
	},
	"night": {
		Title: "Night starts",
		Text:  "Dark enough for astronomical observations",
		Tags:  []string{"sun"},
	},
	"nadir": {
		Title: "Nadir",
		Text:  "Darkest moment of the night, sun is in the lowest position",
		Tags:  []string{"sun"},
	},
	"nightEnd": {
		Title: "Night ends",
		Text:  "Morning astronomical twilight starts",
		Tags:  []string{"sun"},
	},
	"nauticalDawn": {
		Title: "Nautical dawn",
		Text:  "Morning nautical twilight starts",
		Tags:  []string{"sun"},
	},
	"dawn": {
		Title: "Dawn",
		Text:  "Morning nautical twilight ends, morning civil twilight starts",
		Tags:  []string{"sun"},
	},
	"moonrise": {
		Title: "Moonrise",
		Text:  "Top edge of the moon appears on the horizon",
		Tags:  []string{"moon"},
	},
	"moonset": {
		Title: "Moonset",
		Text:  "Moon disappears below the horizon",
		Tags:  []string{"moon"},
	},
	"noon": {
		Title: "Noon",
		Text:  "12 o'clock in the daytime",
		Tags:  []string{"time"},
	},
	"midnight": {
		Title: "Midnight",
		Text:  "12 o'clock in the night",
		Tags:  []string{"time"},
	},
}
