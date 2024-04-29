package main

func maximumWealth(accounts [][]int) int {
    var HighestNumber int

    for _, node := range accounts {
        var totalCount int 
        for i:=0; i < len(node); i++ {
            totalCount += node[i]
        }

        if totalCount > HighestNumber {
            HighestNumber = totalCount
        }
    }

    return HighestNumber
}

func main() {
	var input1 = [[1,5],[7,3],[3,5]]
	maximumWealth(input1)
}