package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Inspect() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var sb strings.Builder

		// первая строка запроса
		sb.WriteString(fmt.Sprintf("%s %s %s\r\n", req.Method, req.RequestURI, req.Proto))

		// заголовки
		for name, values := range req.Header {
			for _, value := range values {
				sb.WriteString(fmt.Sprintf("%s: %s\r\n", name, value))
			}
		}
		sb.WriteString("\r\n")

		// тело, если есть
		if req.Body != nil {
			bodyBytes, err := io.ReadAll(req.Body)
			if err == nil && len(bodyBytes) > 0 {
				sb.Write(bodyBytes)
				sb.WriteString("\r\n")
			} else if err != nil {
				sb.WriteString(fmt.Sprintf("[error reading body: %v]\r\n", err))
			}
		}

		// добавляем разделение между запросами
		sb.WriteString("\r\n===============================\r\n\r\n")

		// открываем файл для добавления
		file, err := os.OpenFile("request_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("Error opening file:", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer file.Close()

		_, err = file.WriteString(sb.String())
		if err != nil {
			log.Println("Error writing to file:", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Println("Request successfully logged.")
		rw.WriteHeader(http.StatusOK)
	}
}
