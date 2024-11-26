package mainmodule1

import (
	"fmt"

	"github.com/kinakoman/submodule/submodule1"
	"github.com/kinakoman/submodule/submodule2"
)

func Mainfunc() {
	fmt.Println(submodule1.Text)
	submodule1.Func1()
	submodule2.Func2()
	submodule2.Func3()
}
