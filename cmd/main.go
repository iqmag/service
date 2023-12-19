package main

import (
	"fmt"
	"my_project/internal/sort_filt_service"
)

//func main() {
//	data, err := struct_service.GetData()
//	if err != nil {
//		log.Fatal("Ошибка при получении данных:", err)
//		return
//	}
//	fmt.Printf("%+v", data)
//}

//______________SERVICE___________________________
//func main() {
//
//	router := mux.NewRouter()
//	router.HandleFunc("/", HandleConnection).Methods(http.MethodGet)
//
//	server := &http.Server{
//		Addr:    "127.0.0.1:8585",
//		Handler: router,
//	}
//
//	log.Println("Сервер запущен на адресе:", server.Addr)
//	err := server.ListenAndServe()
//	if err != nil {
//		log.Fatal("Ошибка сервера:", err)
//	}
//}
//
//// HandleConnection обработчик запросов
//func HandleConnection(w http.ResponseWriter, r *http.Request) {
//
//	//Записываем в ответ строку
//	w.Write([]byte("Hello World!"))
//}

func main() {
	result := sort_filt_service.GetResultData()

	fmt.Println(result)
}
