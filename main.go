package main

import (
	"context"
	"fmt"
	"time"
)

// CONTEXT - это хранение значений и сообщение об ошибках.

func main() {
	ctx := context.Background() // пустой родительский контекст

	//context.TODO() он редкий например для тестов или где не уверены что
	//контекст вам понадобиться.

	//Инструмент для завершения какой-либо задачи:
	// этот контекст отменяется по факту в ручную
	withCancel, cancel := context.WithCancel(ctx)
	fmt.Println(withCancel.Err()) // выводит если есть какая-то ошибка<nil>
	cancel()                      // завершает контекст и выводит ошибку о завершении контекста
	fmt.Println(withCancel.Err()) // повторно выводит ошибку context canceled

	//<nil>
	//context canceled

	// есть контекст к которому можно устанавливать дедлайн:

	withTimeout, cancel := context.WithTimeout(ctx, time.Second*2) //время завершения контекста.
	defer cancel()                                                 // завершение контекста
	fmt.Println(<-withTimeout.Done())                              // получаем канал при вызове метода Done и приходит значение
}
