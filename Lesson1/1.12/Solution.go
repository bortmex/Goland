package main

// 5/16 Массивы и срезы
func main() {

	var workArray = [16]uint8{99, 151, 137, 71, 117, 187, 20, 93, 187, 67, 1, 2, 3, 5, 7, 8}

	var local1 = workArray[workArray[10]]
	var local2 = workArray[workArray[11]]

	workArray[workArray[11]] = local1
	workArray[workArray[10]] = local2

	var local11 = workArray[workArray[11]]
	var local22 = workArray[workArray[12]]

	workArray[workArray[12]] = local11
	workArray[workArray[11]] = local22

	var local111 = workArray[workArray[12]]
	var local222 = workArray[workArray[13]]

	workArray[workArray[13]] = local111
	workArray[workArray[12]] = local222
}
