//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

func TestTakeDamage(t *testing.T) {
	p := NewPlayer("Test", 4, 3 )

	for i := 1; i < 4; i++ {
		health, ok := p.takeDamage(1)

		if !ok {
			t.Errorf("Player should be alive")
		}

		if health != p.MaxHealth - i {
			t.Errorf("Incorrect amount of health: %v, expected %v", health, p.MaxHealth - (i))
		}
	}

	health, ok := p.takeDamage(1)

	if ok {
		t.Errorf("Player should be dead, health is %v", health)
	}
}

func TestReplenishHealth(t *testing.T) {
	p := NewPlayer("Test", 4, 3 )

	p.takeDamage(3)

	for i := 1; i < 4; i++ {
		health := p.replenishHealth(1)

		if health != i + 1 {
			t.Errorf("Incorrect amount of health: %v, expected %v", health, i + 1)
		}
	}

	health := p.replenishEnergy(10)

	if health > p.MaxHealth {
		t.Errorf("Health exceeded max health:  %v, cannot exceed %v", health, p.MaxHealth)
	}
}

func TestUseEnergy(t *testing.T) {
	p := NewPlayer("Test", 4, 3 )

	for i := 1; i < 4; i++ {
		energy, ok := p.useEnergy(1)

		if !ok {
			t.Errorf("Player should be able to use energy")
		}

		if energy != p.MaxEnergy - i {
			t.Errorf("Incorrect amount of energy: %v, expected %v", energy, p.MaxEnergy - (i))
		}
	}

	energy, ok := p.useEnergy(1)

	if ok {
		t.Errorf("Player should have any energy, energy is %v", energy)
	}
}

func TestReplenishEnergy(t *testing.T) {
	p := NewPlayer("Test", 4, 3 )

	p.useEnergy(3)

	for i := 1; i < 4; i++ {
		energy := p.replenishEnergy(1)

		if energy != i  {
			t.Errorf("Incorrect amount of energy: %v, expected %v", energy, i )
		}
	}

	energy := p.replenishEnergy(10)

	if energy > p.MaxEnergy {
		t.Errorf("Energy exceeded max energy:  %v, cannot exceed %v", energy, p.MaxEnergy)
	}
}