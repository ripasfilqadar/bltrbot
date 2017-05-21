package constant

const (
	TOKEN       = "362515117:AAEDHFirPotkKqX63Ffi6PA-aaDPoNCO63M"
	URL_REQUEST = "https://api.telegram.org/bot62515117:AAEDHFirPotkKqX63Ffi6PA-aaDPoNCO63M"
)

func DayOfWeek(day string) string {
	NAME_OF_DAY := map[string]string{
		"Sunday":    "Minggu",
		"Monday":    "Senin",
		"Tuesday":   "Selasa",
		"Wednesday": "Rabu",
		"Thursday":  "Kamis",
		"Friday":    "Jumat",
		"Saturday":  "Sabtu",
	}
	return NAME_OF_DAY[day]
}
