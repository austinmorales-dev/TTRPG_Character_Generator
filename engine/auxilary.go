package engine

import (
	"math/rand"
	"sort"
	"time"

	"github.com/austinmorales/chargen/datastructs"
)

//problem: we need to get an ability score, and return
//the appropriate modifier

func applyModifier(score int) int {
	switch {
	case score == 1:
		return -5
	case score <= 3:
		return -4
	case score <= 5:
		return -3
	case score <= 7:
		return -2
	case score <= 9:
		return -1
	case score <= 11:
		return 0
	case score <= 13:
		return 1
	case score <= 15:
		return 2
	case score <= 17:
		return 3
	case score <= 19:
		return 4
	case score <= 21:
		return 5
	case score <= 23:
		return 6
	case score <= 25:
		return 7
	case score <= 27:
		return 8
	case score <= 29:
		return 9
	case score == 30:
		return 10
	default:
		return 0
	}
}

func DieRoll(sides int, times int) []int {
	rand.Seed(time.Now().UnixNano()) //seed the gen
	rolls := []int{}
	for i := 0; i < times; i++ {
		roll := rand.Intn(sides) + 1
		rolls = append(rolls, roll)
	}
	return rolls
}

func RollAbility() int {
	rolls := DieRoll(6, 4)
	sort.Ints(rolls)
	total := 0
	rolls = rolls[1:]
	for _, v := range rolls {
		total += v
	}
	return total
}

// todo: apply mods for race
func GenerateStats(hp int) *datastructs.Stats {
	// classHD := 8
	stat := &datastructs.Stats{}
	stat.STR = RollAbility()
	stat.STRmod = applyModifier(stat.STR)
	stat.DEX = RollAbility()
	stat.DEXmod = applyModifier(stat.DEX)
	stat.CON = RollAbility()
	stat.CONmod = applyModifier(stat.CON)
	stat.INT = RollAbility()
	stat.INTmod = applyModifier(stat.INT)
	stat.WIS = RollAbility()
	stat.WISmod = applyModifier(stat.WIS)
	stat.CHA = RollAbility()
	stat.CHAmod = applyModifier(stat.CHA)
	stat.HP = hp + applyModifier(stat.CON)
	return stat
}

func ApplyRaceMods(mods []int, genStats *datastructs.Stats) *datastructs.Stats {
	stat := make(map[string]int)
	statnames := []string{"str", "dex", "con", "int", "wis", "cha"}
	for i, v := range mods {
		stat[statnames[i]] = v
	}
	//we've mapped the mods to their appropriate values in the stat map
	//now we need to apply said values to the genStats struct
	genStats.STR += stat[statnames[0]]
	genStats.DEX += stat[statnames[1]]
	genStats.CON += stat[statnames[2]]
	genStats.INT += stat[statnames[3]]
	genStats.WIS += stat[statnames[4]]
	genStats.CHA += stat[statnames[5]]
	return genStats
}
