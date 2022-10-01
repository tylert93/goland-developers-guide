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
	Health, MaxHealth, Energy, MaxEnergy uint
}

func (player *Player) printStats() {
	fmt.Println("\n---- Current Stats for", player.Name, "----")
	fmt.Println("Health:", player.Health)
	fmt.Println("Energy:", player.Energy)
	fmt.Println()
}

func (player *Player) takeDamage(damage uint) {
	player.Health -= damage

	if player.Health <= 0 {
		panic("\nPlayer has died")
	}

	fmt.Println("Damage taken")
	player.printStats()
}

func (player *Player) replenishHealth(health uint) {
	player.Health += health

	if player.Health > player.MaxHealth {
		player.Health = player.MaxHealth
	}

	fmt.Println("Health replenished")
	player.printStats()
}

func (player *Player) useEnergy(energy uint) {
	if energy > player.Energy {
		fmt.Println("Player cannot move")
		return
	}

	player.Energy -= energy

	fmt.Println("Energy used")
	player.printStats()
}

func (player *Player) replenishEnergy(energy uint) {
	player.Energy += energy

	if player.Energy > player.MaxEnergy {
		player.Energy = player.MaxEnergy
	}

	fmt.Println("Energy replenished")
	player.printStats()
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
