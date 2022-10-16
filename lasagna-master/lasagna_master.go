package lasagna

// PreparationTime returns the average preparation time considering the layers of lasagna
func PreparationTime(layers []string, prepTime int) int {
	layersLen := len(layers)

	if prepTime == 0 {
		prepTime = 2
	}

	return layersLen * prepTime
}

// Quantities returns the quantity necessary of noodles and sauce considering the layers of lasgna
func Quantities(layers []string) (int, float64) {
	var noodle int
	var sauce float64

	for _, ingr := range layers {
		if ingr == "noodles" {
			noodle += 50
		}

		if ingr == "sauce" {
			sauce += 0.2
		}
	}

	return noodle, sauce
}

// AddSecretIngredient will add the secret ingredient of friend's recipe on my recipe
func AddSecretIngredient(friendsRecipe []string, myRecipe []string) {
	secretIngr := friendsRecipe[len(friendsRecipe)-1]

	myRecipe[len(myRecipe)-1] = secretIngr
}

// ScaleRecipe must scale the recipe for more portions, returns a slice containing the scaled recipe
func ScaleRecipe(amount []float64, portions int) []float64 {
	var scaled []float64

	for _, a := range amount {
		correctAmount := a * float64(portions) / 2
		scaled = append(scaled, correctAmount)
	}

	return scaled
}
