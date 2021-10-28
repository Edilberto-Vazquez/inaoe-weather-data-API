package etl

type WeatherCloudRow struct {
	Place    string
	DateTime string
	TempIn   string
	Temp     string
	Chill    string
	DewIn    string
	Dew      string
	HeatIn   string
	Heat     string
	Humin    string
	Hum      string
	Wspdhi   string
	Wspdavg  string
	Wdiravg  string
	Bar      string
	Rain     string
	RainRate string
}

func NewWeatherCloudRow(path string, record string) *WeatherCloudRow {
	fields := splitString(record)
	return &WeatherCloudRow{
		Place:    newPlace(path),
		DateTime: fields[0],
		TempIn:   commaToPoint(fields[1]),
		Temp:     commaToPoint(fields[2]),
		Chill:    commaToPoint(fields[3]),
		DewIn:    commaToPoint(fields[4]),
		Dew:      commaToPoint(fields[5]),
		HeatIn:   commaToPoint(fields[6]),
		Heat:     commaToPoint(fields[7]),
		Humin:    commaToPoint(fields[8]),
		Hum:      commaToPoint(fields[9]),
		Wspdhi:   commaToPoint(fields[10]),
		Wspdavg:  commaToPoint(fields[11]),
		Wdiravg:  commaToPoint(fields[12]),
		Bar:      commaToPoint(fields[13]),
		Rain:     commaToPoint(fields[14]),
		RainRate: commaToPoint(fields[15]),
	}
}
