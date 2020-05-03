package utility

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/koni-kuliner/entity"
	"github.com/koni-kuliner/resource/response"
)

func SendSuccessResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	response := response.SuccessResponse{
		Data: data,
		Meta: response.MetaInfo{
			HttpStatus: statusCode,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Meta.HttpStatus)
	_ = json.NewEncoder(w).Encode(&response)
}

func SendSuccessResponseWithLimitAndOffset(w http.ResponseWriter, data interface{}, statusCode int, mapper map[string][]string, count int) {
	limit, _ := strconv.Atoi(mapper["limit"][0])
	offset, _ := strconv.Atoi(mapper["offset"][0])
	response := response.SuccessResponseWithMeta{
		Data: data,
		Meta: response.MetaInfoWithOffsetAndLimit{
			HttpStatus: statusCode,
			Offset:     offset,
			Limit:      limit,
			Total:      count,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Meta.HttpStatus)
	_ = json.NewEncoder(w).Encode(&response)
}

func SendErrorResponse(w http.ResponseWriter, err entity.CustomError) {
	response := response.ErrorResponse{
		Errors: []response.ErrorInfo{
			{
				Message: err.Error(),
				Code:    err.Code,
				Field:   err.Field,
			},
		},
		Meta: response.MetaInfo{
			HttpStatus: err.HttpStatus,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Meta.HttpStatus)
	_ = json.NewEncoder(w).Encode(&response)
}
