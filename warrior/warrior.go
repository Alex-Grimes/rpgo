package warrior

import (
	"math"
	hero "rpgo/hero"
)

type Warrior struct {
	hero.Hero
	Staminas map[string]int
	Stamina  int
	Weapon   string
}

func (w Warrior) Hit() int {
	if w.Stamina >= w.Staminas[w.Weapon] {
		levelEffect := int(math.Ceil(float64(w.Level) * 0.1))
		return w.Attacks[w.Weapon] + levelEffect
	} else {
		return 0
	}
}

func (w *Warrior) TakeDamage(damage int) {
	w.HP = w.HP - damage
}

func (w Warrior) GetInfo() (string, int) {
	return w.Name, w.HP
}

func (w Warrior) GetStamina() int {
	return w.Stamina
}

func (w *Warrior) SpendStamina() {
	if w.Stamina >= w.Staminas[w.Weapon] {
		w.Stamina = w.Stamina - w.Staminas[w.Weapon]
	}
}

func CreateWarrior(hp int, stamina int) *Warrior {
	return &Warrior{
		Hero: hero.Hero{
			Attacks: map[string]int{"Club": 15, "Axe": 30, "Sword": 40},
			HP:      hp,
		},

		Stamina:  stamina,
		Staminas: map[string]int{"Club": 5, "Axe": 10, "Sword": 30},
	}
}

func (w Warrior) IsDown() bool {
	return w.HP <= 0
}
