package validators

import (
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/schemas"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/utils"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var allowedDateparts = []string{
	"day",
	"month",
	"year",
}

var allowedTables = []string{
	"electric_fields",
	"weather_clouds",
}

var allowedFields = map[string][]string{
	"electric_fields": {
		// "lightning",
		// "distance",
		"electric_field",
		// rotor_status
	},
	"weather_clouds": {
		"temp_in",
		"temp",
		"chill",
		"dew_in",
		"dew",
		"heat_in",
		"heat",
		"hum_in",
		"hum",
		"wspdhi",
		"wspdavg",
		"wdiravg",
		"bar",
		"rain",
		"rain_rate",
	},
}

func FieldValidator(sl validator.StructLevel) {
	st := sl.Current().Interface().(schemas.ByDateSchema)
	if ok := utils.FindAllowedValues(st.DatePart, allowedDateparts); !ok {
		sl.ReportError(st.DatePart, "DatePart", "datepart", "alloweddateparts", "")
	}
	if ok := utils.FindAllowedValues(st.Table, allowedTables); !ok {
		sl.ReportError(st.Table, "Table", "table", "allowedtables", "")
	}
	for i := 0; i < len(st.Fields); i++ {
		if ok := utils.FindAllowedValues(st.Fields[i], allowedFields[st.Table]); ok {
			continue
		} else {
			sl.ReportError(st.Fields, "Fields", "fields", "allowedfields", "")
		}
	}
}

func RegisterValidations() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(FieldValidator, schemas.ByDateSchema{})
	}
}
