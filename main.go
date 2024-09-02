package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var mainTemplate = `
<html>
	<body>
		<form action='result'>
			<input name='a'>
			<select name='op'>
				<option value='add'>+</option>
				<option value='sub'>-</option>
				<option value='mul'>*</option>
				<option value='div'>/</option>
			</select>
			<input name='b'>
			<input type='submit'>
		</form>
	</body>
</html>
`

var resultTemplate = `
<html>
	<body>
		<p>Ответ: %v</p>
        <a href='/'>Ещё раз</a>
	</body>
</html>
`

var errorTemplate = `
<html>
    <body>
        <p>Ошибка в данных</p>
        <a href='/'>Назад</a>
    </body>
</html>
`

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/result", resultHandler)
	http.ListenAndServe("localhost:8000", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, mainTemplate)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	urlValues := r.URL.Query()
	a, errA := strconv.Atoi(urlValues.Get("a"))
	b, errB := strconv.Atoi(urlValues.Get("b"))

	if errA != nil || errB != nil {
		fmt.Fprintf(w, errorTemplate)
	} else {
		fmt.Fprintf(w, resultTemplate, apply(a, b, urlValues.Get("op")))
	}
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func apply(a int, b int, op string) int {
	var result int

	switch op {
	case "add":
		result = add(a, b)
	case "sub":
		result = sub(a, b)
	case "mul":
		result = mul(a, b)
	case "div":
		result = div(a, b)
	}

	return result
}
