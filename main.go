package main

func main() {
	var balance uint32 = 1_500_000_000       // 15 миллионов в копейках
	var transfer uint32 = 1_000_000_000      // 10 миллионов в копейках
	total := balance + transfer // uint32 + uint32 будет int32
	println(total)
}
