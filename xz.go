package main
import ("fmt")

func main() {
  var (
    sum1 = 6205 + 7408 // 150 (100 + 50)
    sum2 = sum1 + 8750 // 400 (150 + 250)
    sum3 = sum2 + sum2 // 800 (400 + 400)
  )
  fmt.Println(sum3)
}
