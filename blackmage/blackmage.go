package blackmage

import (
	"math"
	hero "rpgo/hero"
)

type (
	Blackmage struct {
		hero.Hero
		Manas map[string]int
		Mana  int
		Magic string
	}
)

func CreateBlackmage(hp int, mana int) *Blackmage {
	return &Blackmage{
		Hero: hero.Hero{
			Attacks: map[string]int{"Fire1": 30, "Bolt1": 20, "Ice1": 15},
			HP:      hp,
		},
		Mana:  mana,
		Manas: map[string]int{"Fire1": 30, "Bolt1": 10, "Ice1": 5},
	}
}

func (bm Blackmage) Hit() int {
	if bm.Mana >= bm.Manas[bm.Magic] {
		levelEffect := int(math.Ceil(float64(bm.Level) * 0.1))
		return bm.Attacks[bm.Magic] + levelEffect
	} else {
		return 0
	}
}

func (bm Blackmage) TakeDamage(damage int) {
	bm.HP = bm.HP - damage
}

func (bm Blackmage) GetInfo() (string, int) {
	return bm.Name, bm.HP
}

func (bm Blackmage) GetMana() int {
	return bm.Mana
}

func (bm *Blackmage) SpendMana() {
	if bm.Mana >= bm.Manas[bm.Magic] {
		bm.Mana = bm.Mana - bm.Manas[bm.Magic]
	}
}

func (bm Blackmage) IsDown() bool {
	return bm.HP <= 0
}
