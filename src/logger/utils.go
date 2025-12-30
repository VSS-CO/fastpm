package logger

func Prefix(tag, text string) string {
	return "[" + tag + "] " + text
}
