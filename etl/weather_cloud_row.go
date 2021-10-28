package etl

type WeatherCloudRow struct {
	place    string
	dateTime string
	tempIn   string
	temp     string
	chill    string
	dewIn    string
	dew      string
	heatIn   string
	heat     string
	humin    string
	hum      string
	wspdhi   string
	wspdavg  string
	wdiravg  string
	bar      string
	rain     string
	rainRate string
}

func NewWeatherCloudRow(record string, path string) *WeatherCloudRow {
	fields := splitString(record)
	return &WeatherCloudRow{
		place:    newPlace(path),
		dateTime: fields[0],
		tempIn:   commaToPoint(fields[1]),
		temp:     commaToPoint(fields[2]),
		chill:    commaToPoint(fields[3]),
		dewIn:    commaToPoint(fields[4]),
		dew:      commaToPoint(fields[5]),
		heatIn:   commaToPoint(fields[6]),
		heat:     commaToPoint(fields[7]),
		humin:    commaToPoint(fields[8]),
		hum:      commaToPoint(fields[9]),
		wspdhi:   commaToPoint(fields[10]),
		wspdavg:  commaToPoint(fields[11]),
		wdiravg:  commaToPoint(fields[12]),
		bar:      commaToPoint(fields[13]),
		rain:     commaToPoint(fields[14]),
		rainRate: commaToPoint(fields[15]),
	}
}
