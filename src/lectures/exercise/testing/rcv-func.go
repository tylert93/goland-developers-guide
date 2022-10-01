//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	Name string
	Health, MaxHealth, Energy, MaxEnergy int
}

func NewPlayer(name string, health, energy int) Player {
	return Player{name, health, health, energy, energy}
}

func (player *Player) printStats() {
	fmt.Println("\n---- Current Stats for", player.Name, "----")
	fmt.Println("Health:", player.Health)
	fmt.Println("Energy:", player.Energy)
	fmt.Println()
}

func (player *Player) takeDamage(damage int) (int, bool) {
	player.Health -= damage

	if player.Health <= 0 {
		return player.Health, false
	}

	fmt.Println("Damage taken")
	player.printStats()
	return player.Health, true
}

func (player *Player) replenishHealth(health int) int {
	player.Health += health

	if player.Health > player.MaxHealth {
		player.Health = player.MaxHealth
	}

	fmt.Println("Health replenished")
	player.printStats()
	return player.Health
}

func (player *Player) useEnergy(energy int) (int, bool) {
	if energy > player.Energy {
		fmt.Println("Player cannot move")
		return player.Energy, false
	}

	player.Energy -= energy
	fmt.Println("Energy used")
	player.printStats()
	return player.Energy, true
}

func (player *Player) replenishEnergy(energy int) int {
	player.Energy += energy

	if player.Energy > player.MaxEnergy {
		player.Energy = player.MaxEnergy
	}

	fmt.Println("Energy replenished")
	player.printStats()
	return player.Energy
}

func main() {
	playerOne := Player{
		Name: "Tom",
		Health: 100,
		MaxHealth: 100,
		Energy: 50,
		MaxEnergy: 50,
	}

	playerOne.takeDamage(50)

	playerOne.useEnergy(20)

	playerOne.replenishHealth(40)

	playerOne.replenishEnergy(30)
}
