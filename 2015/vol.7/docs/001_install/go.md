----
Go言語のインストールについて
----

# インストール

## 全OS共通版

下記のURLからGoをダウンロードしてください。

https://golang.org/dl/


|OS|ファイル名|
|Windows|go1.4.2.windows-amd64.msi(64bit) / go1.4.2.windows-386.msi(32bit)|
|OSX|go1.4.2.darwin-amd64-osx10.8.pkg|
|Linux|go1.4.2.linux-amd64.tar.gz|

Windows, OSXはインストーラー経由でインストール可能です。
Linux版は解凍するとその中に必要なものが入っているので、
展開先ディレクトリをGOROOTという環境変数に設定して下さい

```sh
$ export GOROOT="/path/to/your go dir"
```

## Windows＠Chocolately

※ VirtualBoxやVagrantに慣れている方は、あえてWindowsを使わずに、Linux上でやった方がやりやすいかもしれません。

ここではWindowsで利用する一例として、
Chocolatelyというパッケージ管理ツールを使用した手順を紹介します。
（Windows10で入る予定のOneGetはうまく動かず断念...）

### Chocolatelyのインストール

コマンドプロンプトの起動（管理者権限）

```
Ctl+x -> 「コマンドプロンプト（管理者）」
```

```sh
C:\Windows\system32 > 

# ↓の行をコピペする コマンドプロンプトでは「Windows左上のアイコンをクリック」→「編集」→「ペースト」で貼り付けできます
@powershell -NoProfile -ExecutionPolicy unrestricted -Command "iex ((new-object net.webclient).DownloadString('https://chocolatey.org/install.ps1'))"  &&  SET PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin
```

### Goのインストール

```sh
C:\Windows\system32 > choco install golang

Chocolatey v0.9.9.5
Installing the following packages:
golang
By installing you accept licenses for the packages.

golang v1.4.2
The package golang wants to run 'chocolateyInstall.ps1'.
Note: If you don't run this script, the installation will fail.
Do you want to run the script?
 1) yes
 2) no
 3) print

1 # ← 1をタイプし、ライセンスに同意の上インストールする

~~~ 以下インストール（省略） ~~~

# パスの追加
C:\Windows\system32 > SET PATH=%PATH%:C:\tools\go\bin

# インストールの確認
C:\Windows\system32 > go version
go version go1.4.2 windows/386

# 環境変数の確認
C:\Windows\system32 > go env
set GOARCH=386
set GOBIN=
set GOCHAR=8
set GOEXE=.exe
set GOHOSTARCH=386
set GOHOSTOS=windows
set GOOS=windows
set GOPATH=
set GORACE=
set GOROOT=C:\tools\go # ← GO本体がインストールされている場所
set GOTOOLDIR=C:\tool
set CC=gcc
set GOGCCFLAGS=-m32 -
set CXX=g++
```

`go get`コマンドで必要になるのでgitもインストールします

```sh
C:\Windows\system32 > choco install git.install -version 1.9.5.20150114
C:\Windows\system32 > choco upgrade git.install
C:\Windows\system32 > refreshenv
```


## Mac OSX

OSXではhomebrewを使ってインストールできます。
ここではhomebrewを使用したインストール手順を説明します。

### gitのインストール

XcodeとCommand Line Toolsをインストールしてください

### homebrewのインストール

```sh
$ ruby -e "$(curl -fsSL https://raw.github.com/mxcl/homebrew/go/install)"
$ brew doctor
$ brew update
```

### golangのインストール

```sh
$ brew install go

# インストールの確認
$ go version
go version go1.4.2 darwin/amd64

# 環境変数の確認
$ go env
```


## CentOS

EPELのレポジトリを使ってyum経由でインストール可能です

```sh
# インストール
$ sudo yum install golang --enablerepo=epel

# インストールの確認
$ go version
go version go1.4.2 linux/amd64

# 環境変数の確認
$ go env
GOARCH="amd64"
GOBIN=""
GOCHAR="6"
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH=""
GORACE=""
GOROOT="/usr/lib/golang" # ← GO本体がインストールされている場所
GOTOOLDIR="/usr/lib/golang/pkg/tool/linux_amd64"
CC="gcc"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0"
CXX="g++"
CGO_ENABLED="1"
```

# GOPATHの設定

GOPATHは簡単にいうと、あなたがGoで開発作業を行いたいディレクトリになります。

他の言語で環境や依存モジュールを切り替えるために、
各ディレクトリのルート階層にbundlerとかvirtualenvの設定やpackage.jsonを置くイメージで、
Goでは環境変数で使用するディレクトリを丸ごと切り替えます。

GOPATH以下は以下の3つのディレクトリが配置されます

```
bin/   # ← Goで作られたアプリケーションの実行ファイル。PATHが通っていれば実行可能になる
pkg/   # ← コンパイル済みのキャッシュファイル。ビルド時に使用される。src以下の内容と乖離してしまった場合は中身は丸ごと削除すると良いかも
src/   # ← go getで取得したリモートパッケージがダウンロードされる。自分が開発するソースコードもここへ配置する。
```

GOPATHは任意の好きな場所に設定して大丈夫です。
HOMEディレクトリをGOPATHにしてもいいですし、
GUI環境の場合はHOME以下に1階層おいてGOPATHにしてもいいんじゃないでしょうか。

以下では一例として `$HOME/repos/gohome` 以下にGOPATHを設定する手順を紹介します

```sh
# Linux / OSXでの設定の例
$ mkdir -p ~/repos/gohome
$ echo "export GOPATH=~/repos/gohome" >> ~/.bashrc
$ echo "export PATH=$PATH:$GOPATH/bin" >> ~/.bashrc
$ source  ~/.bashrc

# Windowsでの設定例
C:\ > echo %USERPROFILE%
C:\ > mkdir %USERPROFILE%\repos\gohome
C:\ > SETX GOPATH %USERPROFILE%\repos\gohome
C:\ > SET PATH=%PATH%;%GOPATH%\bin
```

`go env`を実行して、正常にGOPATHが追加されたか確認して下さい

```sh
$ go env
```

# `go get` でのリモートパッケージ取得

Goの面白い特徴の一つとして、`go get`でgithubやbitbucket上のレポジトリから、
直接リモートパッケージとして利用したいソースコードを取得できるという部分があります。

わざわざRubygemsやnpmに登録しなくても、
githubにレポジトリを置いて`go get`するだけで自動的にmasterブランチがダウンロードできます。
（その反面、利用したいバージョン指定が今のところできないため、
プロダクション環境でGoを使う際には自身で該当のパッケージをダウンロードし、バージョンを固定しておく必要があります。）

まずは `goimports` という便利ツールを入れてみましょう。

```sh
$ go get golang.org/x/tools/cmd/goimports
```

正常に実行されていれば、GOPATH以下に以下のファイルとディレクトリが作成されているはずです。

```sh
$GOPATH/bin/goimports
$GOPATH/src/golang.org/x/tools/cmd/goimports
```

goimportsの確認

```sh
$ goimports -h

usage: goimports [flags] [path ...]
  -d=false: display diffs instead of rewriting files
  -e=false: report all errors (not just the first 10 on different lines)
  -l=false: list files whose formatting differs from goimport's
  -w=false: write result to (source) file instead of stdout
```
