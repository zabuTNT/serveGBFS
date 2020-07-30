package main

type GbfsMain struct {
	LastUpdated int         `json:"last_updated"`
	Ttl         int         `json:"ttl"`
	Version     string      `json:"version"`
	Data        interface{} `json:"data"`
}

type Profile struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

//ROOTS

type Autodiscover struct {
	Feeds []Feed `json:"feeds"`
}
type Versions struct {
	Versions []Version `json:"versions"`
}
type Stations struct {
	Stations []interface{} `json:"stations"`
}

type Bikes struct {
	Bikes []Bike `json:"bikes"`
}
type Rental struct {
	Rental_Hours []Hours `json:"rental_hours"`
}
type Regions struct {
	Regions []Region `json:"regions"`
}
type Calendars struct {
	Calendars []Calendar `json:"calendar"`
}
type PricePlans struct {
	Plans []Fare `json:"plans"`
}
type Alerts struct {
	Alerts []Alert `json:"alerts"`
}

//ELEMENTS

type Feed struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Version struct {
	Version string `json:"version"`
	Url     string `json:"url"`
}

type System struct {
	Id          string `json:"system_id"`
	Language    string `json:"language"`
	Name        string `json:"name"`
	Timezone    string `json:"timezone"`
	Short       string `json:"short_name"`
	Operator    string `json:"operator"`
	Url         string `json:"url"`
	PurchaseUrl string `json:"purchase_url"`
	Phone       string `json:"phone_number"`
	Email       string `json:"email"`
	FeedEmail   string `json:"feed_contact_email"`
	License     string `json:"license_url"`
	RentalApps  App    `json:"rental_apps"`
	StartDate   string `json:"start_date"`
}

type App struct {
	Android AppUri `json:"android"`
	Ios     AppUri `json:"ios"`
}

type AppUri struct {
	Store     string `json:"store_uri"`
	Discovery string `json:"discovery_uri"`
}

type Station struct {
	Id            string   `json:"station_id"`
	Name          string   `json:"name"`
	Lat           float32  `json:"lat"`
	Lon           float32  `json:"lon"`
	Short         string   `json:"short_name"`
	Address       string   `json:"address"`
	CrossStreet   string   `json:"cross_street"`
	RegionId      string   `json:"region_id"`
	PostCode      string   `json:"post_code"`
	RentalUri     Uri      `json:"rental_uris"`
	RentalMethods []string `json:"rental_methods"`
	Capacity      int      `json:"capacity"`
}

type Station_Status struct {
	Id             string `json:"station_id"`
	Availables     int    `json:"num_bikes_available"`
	Disables       int    `json:"num_bikes_disabled"`
	DocksAvailable int    `json:"num_docks_available"`
	DocksDisabled  int    `json:"num_docks_disabled"`
	IsInstalled    bool   `json:"is_installed"`
	IsRenting      bool   `json:"is_renting"`
	IsReturning    bool   `json:"is_returning"`
	LastReport     int    `json:"last_reported"`
}

type Bike struct {
	Id         string  `json:"bike_id"`
	Lat        float32 `json:"lat"`
	Lon        float32 `json:"lon"`
	IsReserved bool    `json:"is_reserved"`
	IsDisabled bool    `json:"is_disabled"`
	Uris       Uri     `json:"rental_uris"`
}

type Uri struct {
	Android string `json:"android"`
	Ios     string `json:"ios"`
	Web     string `json:"web"`
}

type Hours struct {
	User_Types []string `json:"user_types"`
	Days       []string `json:"days"`
	Start      string   `json:"start_time"`
	End        string   `json:"end_time"`
}

type Region struct {
	Id   string `json:"region_id"`
	Name string `json:"name"`
}

type Calendar struct {
	StartMonth int `json:"start_month"`
	StartDay   int `json:"start_day"`
	StartYear  int `json:"start_year"`
	EndMonth   int `json:"end_month"`
	EndDay     int `json:"end_day"`
	EndYear    int `json:"end_year"`
}

type Fare struct {
	Id          string `json:"plan_id"`
	Name        string `json:"name"`
	Currency    string `json:"currency"`
	Price       string `json:"price"`
	Taxable     bool   `json:"is_taxable"`
	Description string `json:"description"`
	Url         string `json:"url"`
}
type MTime struct {
	Start int `json:"start"`
	End   int `json:"end"`
}
type Alert struct {
	Id          string   `json:"alert_id"`
	Type        string   `json:"type"`
	Times       []MTime  `json:"times"`
	Summary     string   `json:"summary"`
	Stations    []string `json:"station_ids"`
	Regions     []string `json:"region_ids"`
	Url         string   `json:"url"`
	Description string   `json:"description"`
	LastUpdate  string   `json:"last_updated"`
}
