package character

import "log"

type Character struct {
    Name string
    Health int
    Level int
    Alive bool
}

func NewCharacter(name string) Character {
    return Character{
        Name: name,
        Health: 1000,
        Level: 1,
        Alive: true,
    }
}

func (c *Character) ReceivesDamage(from Character, damage int) {
    if c.Name == from.Name {
        return
    }

    diffLevels := c.Level - from.Level
    log.Println(diffLevels)
    if diffLevels >= 5 {
            reduce := int(float32(damage) * 0.5)
            damage = damage - reduce
    }

    if diffLevels <= -5 {
            augment := int(float32(damage) * 0.5)
            damage = damage + augment
    }

    c.Health = c.Health - damage
    if c.Health < 0 {
        c.Health = 0
    }

    if c.Health == 0 {
        c.Alive = false
    }
}

func (c *Character) ReceivesHeal(from Character, life int) {
    if c.Name != from.Name {
        return
    }

    if c.Health == 0 {
        return
    }

    c.Health = c.Health + life
    if c.Health > 1000 {
        c.Health = 1000
    }
}

