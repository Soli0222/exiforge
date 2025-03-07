package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	// 元のバージョン値を保存
	originalVersion := Version
	// テスト用のバージョンを設定
	Version = "vTest"
	// テスト終了時に元に戻す
	defer func() { Version = originalVersion }()

	// バージョン出力をキャプチャするためのバッファを用意
	buf := new(bytes.Buffer)
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// テスト終了時に標準出力を元に戻す
	defer func() {
		os.Stdout = originalStdout
	}()

	// versionコマンドを実行
	versionCmd.Run(versionCmd, []string{})

	// パイプを閉じて内容をバッファに移す
	w.Close()
	io.Copy(buf, r)

	output := buf.String()

	// 出力に期待される情報が含まれているか確認
	if output == "" {
		t.Error("Version command output is empty")
	}

	// バージョン情報が含まれているか確認
	if !strings.Contains(output, "vTest") {
		t.Errorf("Version output doesn't contain expected version: %s", output)
	}
}
