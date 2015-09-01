//Webサーバを作成するサンプルコード
package main
import (
    "fmt"
    "html"
    "net/http"
    "os"
)

func main() {

    /*
     * ハンドラ（リクエストに対し呼び出される関数）の登録
     */

     //１つ目のハンドラは/helloで呼び出される
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        //レスポンス内容を出力
        fmt.Fprint(w, "<html><body>Hello</body></html>")
    })

    //2つ目のハンドラは/echoで呼び出される
    http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
        //リクエストパラメータの取得
        val := r.FormValue("say")

        //レスポンス内容を出力
        // ※http.EscapeString関数でHTMLで使用できない文字をエスケープ
        fmt.Fprintf(w, "<html><body>echo:%s</body></html>",
            html.EscapeString(val))
    })

    //ウェブサーバを8080ポートで開始する
    //※得たーが起きない限り、ここから復帰しない
    err := http.ListenAndServe(":8080", nil)

    //エラーチェック
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
