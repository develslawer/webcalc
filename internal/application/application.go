package application

import (
	"encoding/json"
	"errors"
	"github.com/develslawer/webcalc/pkg/calculation"
	"net/http"
)

type Application struct {
	config *Config
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(ErrResponse{ErrInvalidExpression.Error()})
		return
	}

	result, err := calculation.Calc(request.Expression)
	if err != nil {
		if errors.Is(err, calculation.ErrInvalidExpression) || errors.Is(err, calculation.ErrDivisionByZero) ||
			errors.Is(err, calculation.ErrUnknownOperator) || errors.Is(err, calculation.ErrStackOverflow) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrResponse{ErrInvalidExpression.Error()})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrResponse{ErrInternalServer.Error()})
		}

	} else {
		json.NewEncoder(w).Encode(Response{result})
	}
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", LoggerMiddleware(CalcHandler))
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
