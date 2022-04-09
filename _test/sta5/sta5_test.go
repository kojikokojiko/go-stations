package sta5_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"testing"
	"time"
)

func TestStation5(t *testing.T) {
	dbPath := "./temp_test.db"
	if err := os.Setenv("DB_PATH", dbPath); err != nil {
		t.Error("エラーが発生しました", err)
		return
	}

	_, err := procStart(t)
	if err != nil {
		t.Error("エラーが発生しました", err)
		return
	}

	resp, err := http.Get("http://localhost:8080/healthz")
	if err != nil {
		t.Error("エラーが発生しました", err)
		return
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			t.Error(err)
			return
		}
	}()

	want := "{\"message\":\"OK\"}\n"
	got, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("エラーが発生しました", err)
		return
	}

	if string(got) != want {
		t.Errorf("期待していない内容です, got = %s, want = %s\n", got, want)
		return
	}
}

func procStart(t *testing.T) (func() error, error) {
	t.Helper()

	cmd := exec.Command("go", "run", "../../main.go")

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	time.Sleep(2 * time.Second)

	stop := func() error {
		p, err := os.FindProcess(os.Getpid())
		if err != nil {
			return err
		}
		return p.Signal(syscall.SIGTERM)
	}

	return stop, nil
}

// package sta5_test

// import (
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"testing"

// 	"github.com/TechBowl-japan/go-stations/db"
// 	"github.com/TechBowl-japan/go-stations/handler/router"
// )

// func TestStation5(t *testing.T) {
// 	dbPath := "./temp_test.db"
// 	if err := os.Setenv("DB_PATH", dbPath); err != nil {
// 		t.Error("dbPathのセットに失敗しました。", err)
// 		return
// 	}

// 	// t.Cleanup(func() {
// 	// 	if err := os.Remove(dbPath); err != nil {
// 	// 		t.Errorf("テスト用のDBファイルの削除に失敗しました: %v", err)
// 	// 		return
// 	// 	}
// 	// })

// 	todoDB, err := db.NewDB(dbPath)
// 	if err != nil {
// 		t.Error("DBの作成に失敗しました。", err)
// 		return
// 	}
// 	r := router.NewRouter(todoDB)
// 	srv := httptest.NewServer(r)
// 	defer srv.Close()
// 	req, err := http.NewRequest(http.MethodGet, srv.URL+"/healthz", nil)
// 	if err != nil {
// 		t.Error("リクエストの作成に失敗しました。", err)
// 		return
// 	}
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		t.Error("リクエストの送信に失敗しました。", err)
// 		return
// 	}
// 	defer func() {
// 		if err := resp.Body.Close(); err != nil {
// 			t.Error("レスポンスのクローズに失敗しました。", err)
// 			return
// 		}
// 	}()

// 	want := "{\"message\":\"OK\"}\n"
// 	got, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Error("レスポンスの読み込みに失敗しました。", err)
// 		return
// 	}

// 	if string(got) != want {
// 		t.Errorf("期待していない内容です, got = %s, want = %s\n", got, want)
// 		return
// 	}
// }
