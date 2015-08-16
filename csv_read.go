//CSVファイルを読み込むサンプルコード
package main
import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {

    //ファイルのオープン
    file, err := os.OpenFile("test.csv", os.O_RDONLY, 0)

    //エラーチェック
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    //ファイルのクローズ
    defer func() {
        file.Close()
    }()

    //csv.Readerを作成
    r := csv.NewReader(file)

    for {
        //1行ずつ読み込み
        record, err := r.Read()
        if err != nil {
            break
        }

        //読み込んだデータを出力
        fmt.Println(record)
    }
}
