package character_test

import (
	"testing"

	"github.com/carlRondoni/go-basic-env/internal/character"
	"github.com/stretchr/testify/assert"
)

/**
    characters -> health start 1000, start lvl 1, alive true/False
    character deal damage others characters -> subtract health, min 0 and alive False
    character heal if not alive false and cannot max health (lvl 1 1000)
    characters cannot damage himself
    characters only heal himself
    dealing damage -> 5+ lvls than attacker damage -50%
    dealing damage -> 5- lvls than attacker damage +50%
*/
func TestCharacter_New(t *testing.T) {
    character := character.NewCharacter("Javi")

    assert.Equal(t, "Javi", character.Name)
    assert.Equal(t, 1000, character.Health)
    assert.Equal(t, 1, character.Level)
    assert.True(t, character.Alive)
}

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

func TestCharacter_Healed(t *testing.T) {
    character1 := character.NewCharacter("Javi")
    character2 := character.NewCharacter("Carl")

    character1.ReceivesDamage(character2, 500)
    character1.ReceivesHeal(character1, 500)
    assert.Equal(t, 1000, character1.Health)
}

func TestCharacter_CannotBeHealed(t *testing.T) {
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

