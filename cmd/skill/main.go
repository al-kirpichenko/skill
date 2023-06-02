package main

import (
	"fmt"
	"github.com/al-kirpichenko/skill/cmd/config"
	"net/http"
)

// функция main вызывается автоматически при запуске приложения
func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {

	conf := config.NewConfig()
	fmt.Println("Running server on", conf.Host)
	return http.ListenAndServe(conf.Host, http.HandlerFunc(webhook))
}

// функция webhook — обработчик HTTP-запроса
func webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// разрешаем только POST-запросы
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// установим правильный заголовок для типа данных
	w.Header().Set("Content-Type", "application/json")
	// пока установим ответ-заглушку, без проверки ошибок
	_, _ = w.Write([]byte(`
      {
        "response": {
          "text": "Извините, я пока ничего не умею"
        },
        "version": "1.0"
      }
    `))
}
