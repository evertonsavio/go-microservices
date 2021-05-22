package main

import (
	"net/http"
)
/*====================================================================================
GO MAIN
====================================================================================*/
func main()  {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request){

		writer.Write([]byte("Hello World"))
	})

	if err :=http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
/*====================================================================================
	err :=http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
====================================================================================*/