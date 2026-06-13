// Go の plugin パッケージを使用して C++ .so プラグインをロードし、
// 文字列を C++ に送信して、処理結果を受け取るデモプログラム。
package main

import (
	"fmt"
	"os"
	"plugin"
	"strings"
)

const pluginPath = "./myplugin/myplugin.so"

func main() {
    fmt.Println("全引数:", os.Args)
    if len(os.Args) < 2 {
        fmt.Println("error , argment none")
        return
    } 	
    fmt.Println(strings.Repeat("=", 60))
    fmt.Println("  Go ⇄ C++ プラグイン通信デモ")
    fmt.Println("  (plugin パッケージ使用)")
    fmt.Println(strings.Repeat("=", 60))

    // ── プラグインのロード ──
    fmt.Printf("\n📦 プラグインをロード中: %s\n", pluginPath)
    p, err := plugin.Open(pluginPath)
    if err != nil {
        fmt.Fprintf(os.Stderr, "❌ プラグインのロードに失敗: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("✅ プラグインのロード成功")

    var argment = os.Args[1]
    if argment == "add" {
        fmt.Println("#add-start\n")
        ragAddFunc, err := p.Lookup("RagAdd")
        if err != nil {
            fmt.Fprintf(os.Stderr, "❌ RagAdd のルックアップに失敗: %v\n", err)
            os.Exit(1)
        }
        ragAddString, ok := ragAddFunc.(func() string)
        if !ok {
            fmt.Fprintf(os.Stderr, "❌ RagAdd の型が不正\n")
            os.Exit(1)
        }        
        result := ragAddString()
        fmt.Printf("  C++ → Go: %s\n", result)        
        return
    }
    if argment == "search" {
        if len(os.Args) < 3 {
            fmt.Println("error , argment none")
            return
        }
        var input = os.Args[2]
        fmt.Println("title=", input)
        ragSearchFunc, err := p.Lookup("RagSearch")
        if err != nil {
            fmt.Fprintf(os.Stderr, "❌ RagSearch のルックアップに失敗: %v\n", err)
            os.Exit(1)
        }
        ragSearchString, ok := ragSearchFunc.(func(string) string)
        if !ok {
            fmt.Fprintf(os.Stderr, "❌ RagSearch の型が不正\n")
            os.Exit(1)
        }
        result := ragSearchString(input)
        fmt.Printf(" AI: %s\n\n", result)        
        return
    }

    // ── 関数のルックアップ ──
    processStringSym, err := p.Lookup("ProcessString")
    if err != nil {
        fmt.Fprintf(os.Stderr, "❌ ProcessString のルックアップに失敗: %v\n", err)
        os.Exit(1)
    }
    processString, ok := processStringSym.(func(string) string)
    if !ok {
        fmt.Fprintf(os.Stderr, "❌ ProcessString の型が不正\n")
        os.Exit(1)
    }

    toUpperSym, err := p.Lookup("ToUpper")
    if err != nil {
        fmt.Fprintf(os.Stderr, "❌ ToUpper のルックアップに失敗: %v\n", err)
        os.Exit(1)
    }
    toUpper, ok := toUpperSym.(func(string) string)
    if !ok {
        fmt.Fprintf(os.Stderr, "❌ ToUpper の型が不正\n")
        os.Exit(1)
    }

    stringInfoSym, err := p.Lookup("StringInfo")
    if err != nil {
        fmt.Fprintf(os.Stderr, "❌ StringInfo のルックアップに失敗: %v\n", err)
        os.Exit(1)
    }
    stringInfo, ok := stringInfoSym.(func(string) string)
    if !ok {
        fmt.Fprintf(os.Stderr, "❌ StringInfo の型が不正\n")
        os.Exit(1)
    }

    // ── テストメッセージの送受信 ──
    messages := []string{
        "Hello, C++!",
        "こんにちは世界",
        "Go から C++ へ文字列送信テスト",
        "ABCDE 12345",
    }

    // 1. ProcessString（文字列反転）
    fmt.Println("\n" + strings.Repeat("-", 60))
    fmt.Println("📝 テスト1: 文字列処理 (ProcessString)")
    fmt.Println(strings.Repeat("-", 60))
    for _, msg := range messages {
        fmt.Printf("\n  Go → C++: \"%s\"\n", msg)
        result := processString(msg)
        fmt.Printf("  C++ → Go: %s\n", result)
    }

    // 2. ToUpper（大文字変換）
    fmt.Println("\n" + strings.Repeat("-", 60))
    fmt.Println("📝 テスト2: 大文字変換 (ToUpper)")
    fmt.Println(strings.Repeat("-", 60))
    for _, msg := range messages {
        fmt.Printf("\n  Go → C++: \"%s\"\n", msg)
        result := toUpper(msg)
        fmt.Printf("  C++ → Go: %s\n", result)
    }

    // 3. StringInfo（文字列情報）
    fmt.Println("\n" + strings.Repeat("-", 60))
    fmt.Println("📝 テスト3: 文字列情報 (StringInfo)")
    fmt.Println(strings.Repeat("-", 60))
    for _, msg := range messages {
        fmt.Printf("\n  Go → C++: \"%s\"\n", msg)
        result := stringInfo(msg)
        fmt.Printf("  C++ → Go: %s\n", result)
    }

    fmt.Println("\n" + strings.Repeat("=", 60))
    fmt.Println("  ✅ すべてのテスト完了")
    fmt.Println(strings.Repeat("=", 60))
}
