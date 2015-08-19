//指定した文字列の位置を取得するサンプルコード
package main
import (
    "fmt"
    "strings"
)

func main() {

    //文字列内の「Go」の位置（０〜）を取得する
    fmt.Println(strings.Index("The Go Programming Language", "Go"))

    //x見つからないときは−1が返される
    fmt.Println(strings.Index("The Go Programming Language", "xxx"))

}
