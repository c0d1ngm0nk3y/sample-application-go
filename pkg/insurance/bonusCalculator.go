package insurance

func CalculateBonus(costs float64, count int) float64 {
	bonus := 0.0

	if count >= 3 && costs > 350 {
		bonus = bonus + 25
	}

	if count >= 4 && costs >= 600 { //fix
		bonus = bonus + 25
	}

	return bonus
}
