module bosi_tren_poc_go

go 1.22.4

require base_travel_solution v0.0.0

require (
	admissible_offer v0.0.0-00010101000000-000000000000 // indirect
	admissible_service v0.0.0-00010101000000-000000000000 // indirect
	github.com/basgys/goxml2json v1.1.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	travel_solution v0.0.0-00010101000000-000000000000 // indirect
	utils v0.0.0-00010101000000-000000000000 // indirect
)

replace (
	admissible_offer => ./admissible_offer
	admissible_service => ./admissible_service
	base_travel_solution => ./base_travel_solution
	travel_solution => ./travel_solution
	utils => ./utils
)
