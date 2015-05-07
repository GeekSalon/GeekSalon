----
Visual Studio Code用のGOの設定
----


# Visual Studio Code

Visual Studio Codeはマルチプラットフォームで動作する無料のIDEです。

まだプレビューリリースされたばかりなので、機能的に不十分な部分もありますが、
UIはAtom風で、gulpタスクと連携できたりと機能が豊富なIDEです。

Go言語は、シンタックスハイライトと対応するブラケットの表示のみサポートしているようです。
https://code.visualstudio.com/Docs/languages

今のところはユーザーによる言語設定の拡張に対応していないようですが、考えてはいるらしいです。
近い内に対応され、Go関連の機能が満載になることを祈りながら使いましょう。

（私はSublimeText3の動作が重くなってきたので、こちらに切り替えようと思ってます）

## ダウンロード

下記URLからダウンロードしてください

https://code.visualstudio.com/Download

## OSXでのインストール

解凍してでてきた `Visual Studio Code.app` を `アプリケーション` に入れればOKです。


下記の関数を .bashrc に追加しておけば、
`code` コマンドで ファイルやディレクトリをVSCodeで開くことが出来ます。

```sh
code () {
    if [[ $# = 0 ]]
    then
        open -a "Visual Studio Code"
    else
        [[ $1 = /* ]] && F="$1" || F="$PWD/${1#./}"
        open -a "Visual Studio Code" --args "$F"
    fi
}
```

## Linuxでのインストール

適当なディレクトリに配置して解凍してください。
そして`Code`という実行ファイルにPATHを通すため、シンボリックリンクを張ってください

```sh
$ sudo chmod +x /path/to/VSCode/Code
$ sudo ln -s /path/to/VSCode/Code /usr/local/bin/code
```


## Windowsでのインストール

ダウンロードした、 `VSCodeSetup.exe` を実行するとインストールが自動で始まり、完了するとVSCodeが起動します。

インストール先は下記になるようです

```sh
C:\Users\<username>\AppData\Local\Code
```

## 設定

VSCodeでは設定をJSONファイルで行います。（SublimeTextやAtomと同じですね）

OSXでは画面左上の`「Code -> Preference」`から、Windowsでは`「File -> Preference」`から設定可能です。

### 基本設定

共通の設定は、
OSXでは `cmd+,`、 Windowsでは`「File -> Preference -> User Setting」`から設定可能です。

左側ペインがデフォルトの設定、右側がユーザーの個別設定を記述する部分となります。

ここでは一例として、フォント、ホワイトスペースの表示、MarkdownプレビューのCSS設定を行っています。
コメントはつけられるようですが、JSONなので末尾のカンマの有無に気をつけてください。

```json
// Place your settings in this file to overwrite the default settings
{
	// 使用フォントの指定
//	"editor.fontFamily": "Ricty",  // ←はRictyフォントがインストールされていないと表示できません
	// フォントサイズ
	"editor.fontSize": 17, 
	// ホワイトスペースの表示
	"editor.renderWhitespace": true, 
	// Markdownプレビュー時のスタイル設定
	"markdown.styles": ["http://thomasf.github.io/solarized-css/solarized-dark.css"] 
}
```

### プロジェクト固有の設定

現在開いているディレクトリ（ワークスペース）以下でのみ有効な設定は、
OSXでは「Code -> Preference -> Workspace Setting」、 Windowsでは「File -> Preference -> Workspace Setting」から設定可能です。

設定できる内容は、↑の基本設定と同様です。
設定した内容は、 `<ワークスペース>/.setting/setting.json` に保存されます


### ホットキー設定

キーボードショートカットの設定は、
OSXでは「Code -> Preference -> Keyboard Shortcuts」、 Windowsでは「File -> Preference -> Keyboard Shortcuts」から設定可能です。

ここでは一例として、タスク一覧のショートカットを追加しています。

```json
// Place your key bindings in this file to overwrite the defaults
[
	// タスク一覧のショートカット
	{ "key": "alt+cmd+t",  "command": "workbench.action.tasks.runTask" }
]
```

下記にデフォルトのキー設定が記述されています。
便利な機能が記述されているので、英語ですが一度目を通しておくことをおすすめします。

https://code.visualstudio.com/Docs/customization#_default-keyboard-shortcuts

### タスク設定

ワークスペース内で有効なタスクの設定が可能です。

`cmd+shift+P`を押し、表示されたフォームに `tasks` と入力すると、
タスク関連で実行可能なコマンドが表示されます。
ここで `Run Task` を選択すると、実行可能なタスク一覧が表示されます
（先ほどのホットキー設定を行っていれば、`cmd+opt+T`で直接タスク一覧を表示することが可能です）

デフォルトでは何も登録されていないため、`tsc`（TypeScriptのコンパイル実行）のみ表示されます。
（ワークスペース内にgulpが設定されていれば自動でgulpのタスクが認識されます）

タスクの登録は、
`cmd+shift+P`を押し、表示されたフォームに `tasks` と入力し、
`Configure Task Runner` を選択すると設定ファイルを編集することが可能です

下記では一例としてgoimportsを実行するタスクを追加してみましょう。
（複数タスク登録時にはカンマの有無に注意して下さい）

```json
{
	"version": "0.1.0",
	"command": "goimports",
	// ↓ビルドコマンドとして登録する
	"isBuildCommand": true,

	// always=必ず表示 silent=エラー時のみ表示（goimportsでは必ず表示される） never=表示しない
	"showOutput": "silent",
	
	"windows": {
		"command": "goimports"
	},
	
	// 上書きオプションと編集中のファイルを引数に渡す
	"args": ["-w", "${file}"]
}
```

↑では`"isBuildCommand": true`にしているため、ビルドを実行すると、`goimports`が実行されます。
`cmd+shift+B`を押すビルドが実行されます。Outputペインが表示され、自動整形が行われるはずです。
（シンタックスエラーがある場合は、Output内に表示され、goimportsの実行に失敗します）
