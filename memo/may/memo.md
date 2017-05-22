# bit演算

y 0001

^単項で使用するとビット反転     ^y 1110

x 1011

符号ありの型をビット反転させると符号も反転する

```
	var ux uint = 0xff
	var uy uint = 0x1
	x := 0xff
	y := 0x1
	fmt.Printf("int x=%b int y=%b\n", x, y)
	fmt.Printf("^x=%b ^y=%b\n", ^x, ^y)
	fmt.Printf("uint x=%b uint y=%b\n", ux, uy)
	fmt.Printf("^ux=%b ^uy=%b\n", ux, uy)
	fmt.Printf("x&^y=%b \n", x&^y)
	fmt.Printf("y&^x=%b \n", y&^x)
```

&^二項演算子で使用するとビット差   x&^y 1010


# bigの計算
var x,y big.Rat
x.Add(&x,&y) //xを書き換える。

javaの場合bigIntergerに対するaddはGoで言うと下記のようになっている
レシーバー(z)と第一引数(x)が異なるため、加算されるごとにメモリ消費量が増える
したがって、Goはレシーバーと第一引数が一致している場合メモリ効率が良い
x,y,z big.Rat
z.Add(&x,&y)

# UTF-8の自己同期化とは
そもそも同期とは？
例えばCPUアーキテクチャデータバスはバイト列の区切りをどのように判断しているのか
CPUクロックを見てデータの区切りを判断している。つまり、データバスは自己では同期できていない

```イメージ
　　　　　　 0|0|1|0|0|
CPUクロック  ^ ^ ^ ^ ^
```

この意味で、UTF-8はそれ自身でデータの区切りがわかるフォーマットのため自己同期化されていると言える

# Sprintfとstrconv.Itoaどちらを使うべきか
Sprintfは空interfaceを引数に持つ汎用品のため、内部的に
リフレクションを使用しているので、遅い。
たの文字列と組み合わせて使うときはSprintfで良いが、単一の文字列ならItoaで良いのでは


# Goのコンパイル環境
go env goに関する環境変数の表示

[GO_ROOT_BOOTSTRAP Go 1.7.1]を利用して、[~/tools/go/src]にある最新のコンパイラ[仮にNCompとする]が作成される。このNCompを利用して、残りの[~/tools/go/src]ソースコードをコンパイルする。

初期の頃はこのBootStarpコンパイラがC言語で書かれていた。

[~/tools/go/src]内のinternalは外からさわれないパッケージ
vendarパッケージは3rdベンダーのソースコードを丸々取り込んでいるもの
丸々自分たちの領域に取り込むことで、goだけあれば動作することを保証している

# Rune int32
Runeとint32は全く一緒の型

# 配列もしくはArrayへの配列言語仕様書に記載

```go

func zero_array(ptr *[32]byte){
    for i := range ptr {
        ptr[i]=0
    }
}

//コンパイルエラー
func zero_array(ptr []byte) {
    for i := range ptr {
        ptr[i]=0
    }
}

func zero_array(ptr *[]byte) {
    for i := range ptr {
        (*ptr[i])=0
    }
}

```

# 構造体のゼロ値が実用的なデフォルト値ではない場合はどうするのか
newStruct()のようなAPIを提供する

json unmarshalは空interface型を指定するとjsonの型を見て自動的に型変換を行なってくれる