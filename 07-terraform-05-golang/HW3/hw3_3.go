    package main
       
    import "fmt"
    
    func main() {
        var x = []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17,}
        var mi int = 0
		for i := 0; i < len(x); i++ {
			if x[i] < x[mi] {
				mi = i
			}
		}
        fmt.Println(x[mi])
    }
