package character_test

import (
	"testing"

	"github.com/carlRondoni/go-rpg-game/internal/character"
	"github.com/stretchr/testify/assert"
)

// Character
func TestCharacter_New(t *testing.T) {
    character := character.NewCharacter("Javi")

    assert.Equal(t, "Javi", character.Name)
    assert.Equal(t, 1000, character.Health)
    assert.Equal(t, 1, character.Level)
    assert.True(t, character.Alive)
}

// Damage
func TestCharacter_DealDamage(t *testing.T) {
    character1 := character.NewCharacter("Javi")
    character2 := character.NewCharacter("Carl")

    character1.ReceivesDamage(character2, 500)
    assert.Equal(t, 500, character1.Health)
}

func TestCharacter_CannotHurtHimself(t *testing.T) {
    character1 := character.NewCharacter("Javi")

    character1.ReceivesDamage(character1, 500)
    assert.Equal(t, 1000, character1.Health)
}

func TestCharacter_DamageCannotExceedMaxHealth(t *testing.T) {
    character1 := character.NewCharacter("Javi")
    character2 := character.NewCharacter("Carl")

    character1.ReceivesDamage(character2, 1500)
    assert.Equal(t, 0, character1.Health)
    assert.False(t, character1.Alive)
}

func TestCharacter_DeadCharacterCannotReceiveDamage(t *testing.T) {
        character1 := character.NewCharacter("Javi")
        character2 := character.NewCharacter("Carl")

        character1.ReceivesDamage(character2, 1000)
    assert.Equal(t, 0, character1.Health)
    assert.False(t, character1.Alive)

    character1.ReceivesDamage(character2, 1000)
    assert.Equal(t, 0, character1.Health)
    assert.False(t, character1.Alive)
}

func TestCharacter_DamagedByStrongerCharacter(t *testing.T) {
        character1 := character.NewCharacter("Javi")
        character2 := character.NewCharacter("Carl")
        character2.Level = 6

        character1.ReceivesDamage(character2, 500)
    assert.Equal(t, 250, character1.Health)
}
        
func TestCharacter_DamagedByWeakerCharacter(t *testing.T) {
        character1 := character.NewCharacter("Javi")
        character2 := character.NewCharacter("Carl")
        character1.Level = 6

        character1.ReceivesDamage(character2, 500)
    assert.Equal(t, 750, character1.Health)
}

// Heal
func TestCharacter_Healed(t *testing.T) {
    character1 := character.NewCharacter("Javi")
    character2 := character.NewCharacter("Carl")

    character1.ReceivesDamage(character2, 500)
    character1.ReceivesHeal(character1, 500)
    assert.Equal(t, 1000, character1.Health)
}

func TestCharacter_CannotBeHealedByOtherCharacter(t *testing.T) {
    character1 := character.NewCharacter("Javi")
    character2 := character.NewCharacter("Carl")

    character1.ReceivesDamage(character2, 500)
    character1.ReceivesHeal(character2, 500)
    assert.Equal(t, 500, character1.Health)
}

func TestCharacter_CannotHealMoreThanMaxHealth(t *testing.T) {
    character1 := character.NewCharacter("Javi")
    character2 := character.NewCharacter("Carl")

    character1.ReceivesDamage(character2, 500)
    character1.ReceivesHeal(character1, 1500)
    assert.Equal(t, 1000, character1.Health)
}

func TestCharacter_DeadCannotbeHealed(t *testing.T) {
    character1 := character.NewCharacter("Javi")
    character2 := character.NewCharacter("Carl")

    character1.ReceivesDamage(character2, 1000)
    character1.ReceivesHeal(character1, 1000)
    assert.Equal(t, 0, character1.Health)
    assert.False(t, character1.Alive)
}

