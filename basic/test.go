package main

import "fmt"

func main() {
	profiles := map[string]string{"Kampus :": "Universitas Klabat", "Latitude :": "1.4175", "Longitude :": "124.9840"}
	for name, age := range profiles {
		fmt.Println(name, age)
	}
}
