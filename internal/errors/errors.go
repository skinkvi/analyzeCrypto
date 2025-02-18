package errors

import "fmt"

type CustomError struct {
	Package string
	Func    string
	Desc    string
	Err     error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Пакет: %s, Функция: %s, Описание:%s, Ошибка: %v", e.Package, e.Func, e.Desc, e.Err)
}

func Wrap(err error, pkg, funcName, desc string) *CustomError {
	return &CustomError{
		Package: pkg,
		Func:    funcName,
		Err:     err,
		Desc:    desc,
	}
}
