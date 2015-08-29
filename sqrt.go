//平方根を計算するサンプルコード
package main
import (
    "fmt"
    "math"
)

func main() {

    //１から５までの平方根を出力
    for i := 1; i <= 5; i++ {

        //sqrtのパラメータはfloat64型
        fmt.Println(i, math.Sqrt(float64(i)))
    }
}
