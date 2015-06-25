package rtf

import (
	"github.com/revel/revel"
	"time"
	"io"
	"crypto/md5"
	"strings"
	"html/template"
	"fmt"
	"math/rand"
)

func init() {

	revel.TemplateFuncs["nil"] = func(a interface{}) bool {
		return a == nil
	}

	/**
	Number functions
	 */
	revel.TemplateFuncs["add"] = func(a, b int) int {
		return a + b
	}

	revel.TemplateFuncs["ueq"] = func(a, b interface{}) bool {
		return !(a == b)
	}

	revel.TemplateFuncs["minus"] = func(a, b int) int {
		return a - b
	}

	revel.TemplateFuncs["multiply"] = func(args ...int) int {
		result := 0
		for i, value := range args {
			if i == 0 {
				result = value
			} else {
				result *= value
			}
		}
		return result
	}

	revel.TemplateFuncs["divide"] = func(args ...int) int {
		if len(args) < 2{
			return 0
		}
		result := args[0]
		for _, value := range args[1:] {
			if value == 0 {
				return 0
			} else {
				result /= value
			}
		}
		return result
	}

	revel.TemplateFuncs["less"] = func(a, b int) bool {
		return a < b
	}

	revel.TemplateFuncs["lte"] = func(a, b int) bool {
		return a <= b
	}
	revel.TemplateFuncs["gte"] = func(a, b int) bool {
		return a >= b
	}

	revel.TemplateFuncs["percentageToInt"] = func(f float32) int {
		return int(f * 100)
	}

	// perform the percent
	revel.TemplateFuncs["percent"] = func(total int64, current int64) int64 {
		val := int64((float64(current) / float64(total)) * 100)
		if val < 0 {
			return 0
		}
		return val
	}

	/**
	Date and time functions
	 */

	revel.TemplateFuncs["timestamp"] = func() int64 {
		return time.Now().Unix()
	}

	revel.TemplateFuncs["year"] = func() int {
		return time.Now().Year()
	}

	revel.TemplateFuncs["month"] = func() string {
		return time.Now().Month().String()
	}

	revel.TemplateFuncs["week"] = func() string {
		return time.Now().Weekday().String()
	}

	revel.TemplateFuncs["day"] = func() int {
		return time.Now().Day()
	}

	revel.TemplateFuncs["today"] = func() string {
		t := time.Now()
		return t.Format("2006-01-02")
	}

	revel.TemplateFuncs["now"] = func() string {
		t := time.Now()
		return t.Format("2006-01-02 15:04")
	}

	revel.TemplateFuncs["formatTime"] = func(date time.Time, format string) string {
		return date.Format(format)
	}

	revel.TemplateFuncs["dayLeft"] = func(date time.Time) int {
		return int(date.Sub(time.Now()).Hours() / 24)
	}

	/*String functions*/

	revel.TemplateFuncs["lower"] = func(s string) string {
		return strings.ToLower(s)
	}

	revel.TemplateFuncs["upper"] = func(s string) string {
		return strings.ToUpper(s)
	}

	revel.TemplateFuncs["splitComma"] = func(str string) []string {
		return strings.Split(str, ",")
	}

	revel.TemplateFuncs["shortContent"] = func(cont string,length int) string {
		chars := []rune(cont)

		if len(chars) > length {
			return string(chars[:length])+"..."
		}
		return cont;
	}

	revel.TemplateFuncs["join"] = func(ss []string, sept string) string {
		return strings.Join(ss, sept)
	}

	revel.TemplateFuncs["contains"] = func(s, substr string) bool {
		return strings.Contains(s, substr)
	}

	revel.TemplateFuncs["replace"] = func(s string, old string, new string) string {
		return strings.Replace(s, old, new, -1)
	}

	revel.TemplateFuncs["md5"] = func(str string) string {
		h := md5.New()
		io.WriteString(h, str)
		return fmt.Sprintf("%x", h.Sum(nil))
	}

	revel.TemplateFuncs["raw"] = func(str string) template.HTML {
			return template.HTML(str)
	}

	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var alphas = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var numbers = []rune("0123456789")

	revel.TemplateFuncs["randomAlnum"] = func(length int) string {
		rand.Seed(time.Now().UnixNano())
		b := make([]rune, length)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return string(b)
	}

	revel.TemplateFuncs["randomAlphas"] = func(length int) string {
		rand.Seed(time.Now().UnixNano())
		b := make([]rune, length)
		for i := range b {
			b[i] = alphas[rand.Intn(len(alphas))]
		}
		return string(b)
	}

	revel.TemplateFuncs["randomNumeric"] = func(length int) string {
		rand.Seed(time.Now().UnixNano())
		b := make([]rune, length)
		for i := range b {
			b[i] = numbers[rand.Intn(len(numbers))]
		}
		return string(b)
	}
}