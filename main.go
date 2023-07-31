package main

import (
	"wc3_game_tracker/server"
)

// @title           Warcraft III - Game Tracker
// @version         1.0
// @description     A simple and sample API to make WC3 Leagues data more accessible
// @termsOfService  https://flutter-explained.dev/impressum

// @contact.name   Max Weber
// @contact.url    https://flutter-explained.dev/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/v1
func main() {
	server.Setup()
}
