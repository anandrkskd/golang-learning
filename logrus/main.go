package main
import (
	"github.com/sirupsen/logrus"
)
type TerminalFormatter struct{}
func (f *TerminalFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return append([]byte(entry.Message), '\n'), nil
}
func main() {
	logrus.SetFormatter(new(TerminalFormatter))
	logrus.Infof("hello world")
}
