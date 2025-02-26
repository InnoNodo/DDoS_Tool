package main

import (
	"ddos_tool/tools/appearance"
	"ddos_tool/tools/attack"
)

func main() {
	//Use this function to check your proxies
	//check.CheckProxies()

	//Perform attack
	URL := appearance.Banner()

	// You should have file http.txt in root directory for proper work
	attack.PerformAttack(URL)
}
