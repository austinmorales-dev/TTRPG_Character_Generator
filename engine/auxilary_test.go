package engine

import (
	"testing"

	"github.com/austinmorales/chargen/datastructs"
	"github.com/stretchr/testify/assert"
)

func TestModifier(t *testing.T) {
	assert.Equal(t, applyModifier(1), -5)
	assert.Equal(t, applyModifier(10), 0)
	assert.Equal(t, applyModifier(12), 1)
	assert.Equal(t, applyModifier(18), 4)
	assert.Equal(t, applyModifier(30), 10)
	assert.Equal(t, applyModifier(6), -2)
}

func TestRoll(t *testing.T) {
	assert.Equal(t, len(DieRoll(5, 5)), 5)
}

func TestFullStats(t *testing.T) {
	// var cases []datastructs.Stats

	stats := datastructs.Stats{
		HP:  8,
		STR: 8,
		DEX: 8,
		CON: 8,
		INT: 8,
		WIS: 8,
		CHA: 8,
	}
	mods := []int{0, 1, 0, 0, 0, 1}

	rm_stats := ApplyRaceMods(mods, &stats)

	assert.Equal(t, rm_stats, &datastructs.Stats{
		HP:  8,
		STR: 8,
		DEX: 9,
		CON: 8,
		INT: 8,
		WIS: 8,
		CHA: 9,
	})
	assert.NotEqual(t, rm_stats, &datastructs.Stats{
		HP:  9,
		STR: 8,
		DEX: 9,
		CON: 9,
		INT: 9,
		WIS: 9,
		CHA: 9,
	})
}
