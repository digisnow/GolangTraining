//プログラムを指定した終了コードで終了するサンプルコード２
package main
import (
    "log"
    "os"
)

func main() {

    //存在しないファイルのオープン
    _, err := os.OpenFile("test", os.O_RDONLY, 0)
    if err != nil {
        log.Fatalln("エラー", err)
    }
}

