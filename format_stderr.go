//標準エラー出力に値を出力する
package main
import (
    "fmt"
    "os"
)

func main() {

    //値を標準エラー出力に出力
    fmt.Fprint(os.Stderr, "北海道", "の降水確率は", 100, "%です。\n")

    //値を標準エラー出力に出力
    //改行は自動出力されるが、値の間にスペースが出力される
    fmt.Fprintln(os.Stderr, "北海道", "の降水確率は", 100, "%です。")

    //書式を使用して標準エラー出力に出力
    fmt.Fprintf(os.Stderr, "%sの降水確率は%d%%です。\n", "北海道", 100)
}
