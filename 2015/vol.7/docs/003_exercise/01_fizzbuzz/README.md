----
課題 FizzBuzz
----

# 1. 関数 PrintFizzBuzz を作りましょう

## 課題

次の要件を満たす関数 PrintFizzBuzz を作って下さい


- 要件
    - 引数には int を一つ取ります
    - 1から引数として与えられた数までの数値を順番に出力してください
    - ただし3の倍数のときは、数の代わりに「Fizz」を出力して下さい
    - 5の倍数のときは、数の代わりに「Buzz」を出力して下さい
    - 3と5の倍数のときは「FizzBuzz」を出力して下さい
    - それぞれの値を出力した後には改行を付けてください
    - 出力にはfmtを使用して下さい
    - package名には「fizzbuzz」を使用して下さい

- コードのサンプル

```go
// 例
package fizzbuzz

import "fmt"

func PrintFizzBuzz(num int) {
    // ...
}
```

## 出力例

```
func main(){
    PrintFizzBuzz(10)
}
```

```
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
```

## テスト

このGithubのディレクトリに `fizzbuzz_test.go` というテストコードがあるので、
こちらをチェックアウトして、実装コードを同じ場所に配置してください。

また、テストコードをコピーして実装コードと同じディレクトリに配置してもOKです。

```sh
$ cd /path/to/yourcode/  # ← 実装コードが置いてあるディレクトリへ移動する
$ go test ./... -run="PrintFizzBuzz"

ok  GeekSalon/2015/vol.7/docs/003_exercise/01	0.006s
```

もしOKが出ない場合はエラーメッセージを確認してみましょう

----

# 2. 関数 FizzBuzz を作りましょう

## 課題

次の要件を満たす関数 FizzBuzz を作って下さい

- 要件
    - 引数には int を一つ取ります
    - 1から引数として与えられた数までの数値のリストを返却してください
    - ただし3の倍数のときは、数の代わりに「Fizz」という文字列にして下さい
    - 5の倍数のときは、数の代わりに「Buzz」という文字列にして下さい
    - 3と5の倍数のときは「FizzBuzz」という文字列にして下さい
    - package名には「fizzbuzz」を使用して下さい
    - 戻り値には []interface{} （interface{}のスライス）を使用して下さい

```go
// 例
func FizzBuzz(num int) []interface{} {
    // ...
}
```

## テスト

このGithubのディレクトリに `fizzbuzz_test.go` というテストコードがあるので、
こちらをチェックアウトして、実装コードを同じ場所に配置してください。

また、テストコードをコピーして実装コードと同じディレクトリに配置してもOKです。

```sh
$ cd /path/to/yourcode/  # ← 実装コードが置いてあるディレクトリへ移動する
$ go test ./... -run="FizzBuzz"

ok  GeekSalon/2015/vol.7/docs/003_exercise/01	0.007s
```

もしOKが出ない場合はエラーメッセージを確認してみましょう

----

# 3. 関数 Fibonacci を作りましょう

## 課題 a

次の要件を満たす関数 Fibonacci を作って下さい

- 要件
    - 引数には int を一つ取ります
    - 1から引数として与えられた数までの[フィボナッチ数列](http://ja.wikipedia.org/wiki/%E3%83%95%E3%82%A3%E3%83%9C%E3%83%8A%E3%83%83%E3%83%81%E6%95%B0)のリストを返却してください
    - 戻り値には []int （intのスライス）を使用して下さい

```go
// 例
func Fibonacci(num int) []int {
    // ...
}
```

## 課題 b

課題 a で作成した関数 `Fibonacci` をテストするコードを書いて下さい

```go
// 例: fibonacci_test.go
import "testing"

func TestFibonacci(t *testing.T) {
    result := Fibonacci(100)
    // ...以下テスト処理
}
```

作成し終わったら、テストをしてみましょう

```go
$ go test ./...
```

※ 同じディレクトリには同一のパッケージ名しか配置することが出来ません。
fizzbuzzのコードと混在している場合は、ディレクトリを分けるかパッケージ名を統一して下さい


## 課題 c

※ 解法が思い浮かばない場合、この課題はスキップしてください

- 要件
    - 課題 a で作成した関数 `Fibonacci` と同様の要件を満たす関数を作成して下さい
    - 今回作る関数では、より効率の良い（≒計算量が少ない）関数、または効率の悪い関数を作成して下さい
    - 作成したら、既存の関数 `Fibonacci` と新しい関数でベンチマークを取るコードを作成して下さい
    

```go
// ベンチマークの例: fibonacci_test.go

var smallNum = 1000
var bigNum = 200000

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(smallNum)
	}
}

func BenchmarkFibonacciAlpha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciAlpha(smallNum) // 新しく作成した関数
	}
}
```

```sh
$ go test -bench . -benchmem
```
