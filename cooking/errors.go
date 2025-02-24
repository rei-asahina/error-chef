package cooking

import (
	"fmt"
	"runtime"
	"strings"
)

// Error はエラーに追加情報（メッセージ、元エラー、スタックトレース）を付加するための型です。
type Error struct {
	Msg        string // エラーに付加する任意のメッセージ
	Err        error  // 元のエラー
	StackTrace string // エラー発生時のスタックトレース
}

// Error メソッドにより、Error型はerrorインターフェースを満たします。
func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v\n%s", e.Msg, e.Err, e.StackTrace)
	}
	return fmt.Sprintf("%s\n%s", e.Msg, e.StackTrace)
}

// Wrap 関数は、渡されたエラーをラッピングしてスタックトレースを追加します。
func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	return &Error{
		Msg:        msg,
		Err:        err,
		StackTrace: getStackTrace(),
	}
}

// getStackTrace は runtime.Callers を利用してスタックトレースを取得します。
func getStackTrace() string {
	// この関数自身とラッピング関数分のフレームをスキップする
	const skip = 3
	pcs := make([]uintptr, 10)
	n := runtime.Callers(skip, pcs)
	pcs = pcs[:n]

	var strBuilder strings.Builder
	for _, pc := range pcs {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line := fn.FileLine(pc)
		strBuilder.WriteString(fmt.Sprintf("%s:%d %s\n", file, line, fn.Name()))
	}
	return strBuilder.String()
}
