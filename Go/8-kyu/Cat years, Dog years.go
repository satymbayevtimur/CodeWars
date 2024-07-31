package kata

func CalculateYears(humanYears int) [3]int {
    catYears := 0
    dogYears := 0

    if humanYears == 1 {
        catYears = 15
        dogYears = 15
    } else if humanYears == 2 {
        catYears = 15 + 9
        dogYears = 15 + 9
    } else if humanYears >= 3 {
        catYears = 15 + 9 + 4*(humanYears-2)
        dogYears = 15 + 9 + 5*(humanYears-2)
    }

    return [3]int{humanYears, catYears, dogYears}
}
