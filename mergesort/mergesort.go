package mergesort

//Teilen (Divide): Die unsortierte Liste wird rekursiv in zwei Hälften geteilt, bis nur
//  noch Ein-Element-Listen übrig sind. Eine Liste mit nur einem Element gilt per Definition
//  als sortiert.

//Herrschen (Conquer): Die Ein-Element-Listen werden als sortierte Unterlisten betrachtet.

//Kombinieren (Combine): Die sortierten Unterlisten werden paarweise durch die Funktion Merge zu
// größeren, ebenfalls sortierten Listen zusammengeführt, bis die gesamte Liste wieder vereint ist.

// Merge kombiniert zwei sortierte Listen zu einer sortierten Liste.
func Merge(a1 []int, a2 []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(a1) && j < len(a2) {
		if a1[i] <= a2[j] {
			result = append(result, a1[i])
			i++
		} else {
			result = append(result, a2[j])
			j++
		}
	}

	result = append(result, a1[i:]...)
	result = append(result, a2[j:]...)

	return result
}

// MergeSort sortiert die übergebene Liste mittels des Merge-Sort-Algorithmus.
func MergeSort(arr []int) []int {
	// Basisfall: Wenn die Liste weniger als 2 Elemente hat, ist sie bereits sortiert.
	if len(arr) < 2 {
		return arr
	}

	// Teile die Liste in zwei Hälften, sortiere beide Hälften rekursiv und
	// führe die beiden sortierten Hälften mit Merge zusammen.
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return Merge(left, right)
}
