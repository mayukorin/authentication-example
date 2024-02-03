# 必要なもの
- node v20.11.0
- npm v10.2.4
- go v1.21.6

# clone したらやること
1. frontend ディレクトリで，`npm install`
2. frontend ディレクトリで，`node app.js`
3. http://localhost:3000/ で以下のような画面が表示されれば OK.
4. backend ディレクトリで，`go run main.go`

# Basic Authentication
1. Basic Authentication ボタンをクリックする
2. 上に表示されるダイアログに，username: mayukorin, password: password を入力し，OK を押す．
3. Basic Authentication が完了し，Basic Authentication ボタンの横に「hello world!!」が表示される．