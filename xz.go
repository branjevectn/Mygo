package main
import ("fmt")

func main() {
  var (
    sum1 = 46 + 53 // 150 (100 + 50)
    sum2 = sum1 + 66 // 400 (150 + 250)
    sum3 = sum2 + sum2 // 800 (400 + 400)
  )
  fmt.Println(sum3)
}
