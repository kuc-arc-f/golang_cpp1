# go_rag_1

 Version: 0.9.1

 Author  :

 date    : 2026/06/07

 update :

***
GoLang C++ call  , RAG SQLite

* Linux
* go version go1.26.1 linux/amd64
* gcc
* embedding : Gemini-embedding-001
* model: gemma-4-E2B
* llama.cpp , llama-server

***

* LIB add
```
sudo apt-get install libsqlite3-dev
sudo apt install nlohmann-json3-dev
```

***
* llama-server start
* port 8090: gemma-4-E2B

```
#gemma-4-E2B

/usr/local/llama-b8642/llama-server -m /var/lm_data/unsloth/gemma-4-E2B-it-Q4_K_S.gguf \
 --chat-template-kwargs '{"enable_thinking": false}' --port 8090 
```
***
* build

* GEMINI_API_KEY SET
```
export GEMINI_API_KEY=
```

```
sqlite3 ./example.db < table.sql
```

***
```
go mod init example.com/cgo-rag-1
go mod tidy
```

* C++ build
```
make plugin
```

* go buils
```
go build -o app main.go
```
***

* vector data add
```
./app add
```

* search
```
./app search hello
```

***
### blog

https://zenn.dev/knaka0209/scraps/3d848ac18a27a0

***

