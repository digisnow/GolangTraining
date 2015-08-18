//URLエンコード・デコードを行うサンプルコード
package main
import (
    "fmt"
    "net/url"
)

func main() {

    //データ
    data := "abcあいう/:"

    //URLエンコード
    enc := url.QueryEscape(data)

    //エンコード結果を出力
    fmt.Println(enc)

    //エンコード結果をBASE64デコード
    dec, _ := url.QueryUnescape(enc)

    //デコード結果を出力
    fmt.Printf(dec)
}
