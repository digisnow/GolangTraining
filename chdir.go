//カレントディレクトリを変更するサンプルコード
package main
import (
    "os"
    "fmt"
)

func main() {

    //カレントディレクトリを取得して出力
    current, _ := os.Getwd()
    fmt.Println(current)

    //カレントディレクトリを親ディレクトリに変更
    os.Chdir("..")

    //再びカレントディレクトリを出力
    current, _ = os.Getwd()
    fmt.Println(current)
}
