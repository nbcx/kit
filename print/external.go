package print

import (
	"encoding/json"
	"fmt"
)

func Ln(a ...any) {
	_, _ = LnE(a...)
}

func LnE(a ...any) (n int, errno error) {
	return fmt.Println(wrap(a, true)...)
}

// Sprintf is a convenience wrapper for fmt.Sprintf.
//
// Calling Sprintf(f, x, y) is equivalent to
// fmt.Sprintf(f, Formatter(x), Formatter(y)).
func SF(format string, a ...any) string {
	return fmt.Sprintf(format, wrap(a, false)...)
}

// Json 格式化输出 JSON
func Json(a interface{}) {
	_ = JsonE(a)
}

func JsonE(a interface{}) error {
	b, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
