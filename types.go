package heweather

const (
	UnitMetric   = "m"
	UnitImperial = "i"
)

type WeatherAPI struct {
	Username string
	Location string
	Key      string
	Unit     string
	Language string
}

// Basic is the basic information of a place
type Basic struct {
	Location   string  `json:"location"`    //	地区／城市名称
	CityID     string  `json:"cid"`         // 地区／城市ID
	Latitude   float64 `json:"lat,string"`  // 地区／城市纬度
	Longitude  float64 `json:"lon,string"`  // 地区／城市经度
	ParentCity string  `json:"parent_city"` // 该地区／城市的上级城市
	AdminArea  string  `json:"admin_area"`  // 该地区／城市所属行政区域
	Country    string  `json:"cnty"`        // 该地区／城市所属国家名称
	TimeZone   string  `json:"tz"`          // 该地区／城市所在时区
}

// Update contains the update time of data
type Update struct {
	LocalTime string `json:"loc"` // 当地时间，24小时制，格式2017-10-25 12:34
	UTCTime   string `json:"utc"` // UTC时间，24小时制，格式2017-10-25 04:34

}

type DailyForecast struct {
	Date                       string  `json:"date"`               // 预报日期 2013-12-30
	SunRise                    string  `json:"sr"`                 // 日出时间 07:36
	SunSet                     string  `json:"ss"`                 // 日落时间 16:58
	MoonRise                   string  `json:"mr"`                 // 月升时间 04:47
	MoonSet                    string  `json:"ms"`                 // 月落时间 14:59
	TemperatureMax             int     `json:"tmp_max,string"`     // 最高温度
	TemperatureMin             int     `json:"tmp_min,string"`     // 最低温度
	ConditionCodeDay           int     `json:"cond_code_d,string"` // 白天天气状况代码 100
	ConditionCodeNight         int     `json:"cond_code_n,string"` // 晚间天气状况代码 100
	ConditionTextDay           string  `json:"cond_txt_d"`         // 白天天气状况描述 晴
	ConditionTextNight         string  `json:"cond_txt_n"`         // 晚间天气状况描述 晴
	WindDegree                 int     `json:"wind_deg,string"`    // 风向360角度 310
	WindDirection              string  `json:"wind_dir"`           // 风向 西北风
	WindScale                  string  `json:"wind_sc"`            // 风力 1-2
	WindSpeed                  int     `json:"wind_spd,string"`    // 风速，公里/小时 14
	Humidity                   int     `json:"hum,string"`         // 相对湿度 37
	Precipitation              float64 `json:"pcpn,string"`        // 降水量 0
	ProbabilityOfPrecipitation int     `json:"pop,string"`         // 降水概率 0
	Pressure                   int     `json:"pres,string"`        // 大气压强 1018
	UVIndex                    int     `json:"uv_index,string"`    // 紫外线强度指数 3
	Visibility                 int     `json:"vis,string"`         // 能见度，单位：公里 10
}

type RealTimeWeather struct {
	Feel          int     `json:"fl,string"`      // 体感温度，默认单位：摄氏度 23
	Temperature   int     `json:"tmp,string"`       // 	温度，默认单位：摄氏度 21
	ConditionCode int     `json:"cond_code,string"` // 实况天气状况代码 200
	ConditionText string  `json:"cond_txt"`         // 实况天气状况描述 晴
	WindDegree    int     `json:"wind_deg,string"`  // 风向360角度 310
	WindDirection string  `json:"wind_dir"`         // 风向 西北
	WindScale     string  `json:"wind_sc"`          // 风力 1-2
	WindSpeed     int     `json:"wind_spd,string"`  // 风速，公里/小时 14
	Humidity      int     `json:"hum,string"`       // 相对湿度 37
	Precipitation float64 `json:"pcpn,string"`      // 降水量 0
	Pressure      int     `json:"pres,string"`      // 大气压强 1018
	Visibility    int     `json:"vis,string"`       // 能见度，单位：公里 10
	Cloud         int     `json:"cloud,string"`     // 云量 23
}

type HourlyWeather struct {
	Time                string  `json:"time"`             // 预报时间，格式2013-12-30 13:00
	Temperature         int     `json:"tmp,string"`       // 温度，默认单位：摄氏度 21
	ConditionCode       int     `json:"cond_code,string"` // 天气状况代码 101
	ConditionText       string  `json:"cond_txt"`         // 天气状况描述 多云
	WindDegree          int     `json:"wind_deg,string"`  // 风向360角度 310
	WindDirection       string  `json:"wind_dir"`         // 风向 西北
	WindScale           string  `json:"wind_sc"`          // 风力 3-4
	WindSpeed           int     `json:"wind_spd,string"`  // 风速，公里/小时 14
	Humidity            int     `json:"hum,string"`       // 相对湿度 37
	Precipitation       float64 `json:"pcpn,string"`      // 降水量 0
	Pressure            int     `json:"pres,string"`      // 大气压强 1018
	Visibility          int     `json:"vis,string"`       // 能见度，单位：公里 10
	DewPointTemperature int     `json:"dew,string"`       // 露点温度 5
	Cloud               int     `json:"cloud,string"`     // 云量 23
}

type APIResponseData struct {
	Basic           Basic           `json:"basic"`
	DailyForecast   []DailyForecast `json:"daily_forecast"`
	RealTimeWeather RealTimeWeather `json:"now"`
	HourlyWeather   []HourlyWeather `json:"hourly"`
	Status          string          `json:"status"`
	Update          Update          `json:"update"`
}

type APIResponse struct {
	HeWeather6 []APIResponseData `json:"HeWeather6"`
}
