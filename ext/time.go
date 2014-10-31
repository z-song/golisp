package ext

import (
	//"fmt"
	"github.com/z-song/golisp/compiler"
	"time"
)

// (now)
func Now(args []compiler.Node, Env compiler.Environment) (ret interface{}) {

	return compiler.NewNode(time.Now().String())
}

// (year)
func Year(args []compiler.Node, Env compiler.Environment) (ret interface{}) {

	return compiler.NewNode(time.Now().Year())
}

// (month)
func Month(args []compiler.Node, Env compiler.Environment) (ret interface{}) {

	return compiler.NewNode(time.Now().Month())
}

// (day)
func Day(args []compiler.Node, Env compiler.Environment) (ret interface{}) {

	return compiler.NewNode(time.Now().Day())
}

// (hour)
func Hour(args []compiler.Node, Env compiler.Environment) (ret interface{}) {

	return compiler.NewNode(time.Now().Hour())
}

// (minute)
func Minute(args []compiler.Node, Env compiler.Environment) (ret interface{}) {

	return compiler.NewNode(time.Now().Minute())
}

// (second)
func Second(args []compiler.Node, Env compiler.Environment) (ret interface{}) {

	return compiler.NewNode(time.Now().Second())
}

// (yearday)
func YearDay(args []compiler.Node, Env compiler.Environment) (ret interface{}) {

	return compiler.NewNode(time.Now().YearDay())
}

// (timestamp)
func Timestamp(args []compiler.Node, Env compiler.Environment) (ret interface{}) {

	return compiler.NewNode(int(time.Now().Unix()))
}
