package main

import "net/http"


func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	})
	
	// http.ListenAndServe(":8080", nil)
	
	//Error handling
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
			panic(err)
	}

	//Error handling another way error scoping to condition only
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 		panic(err)
	// }
}
