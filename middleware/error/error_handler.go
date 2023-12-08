package error

import (
	"GinBoilerplate/shared"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

type ReturnResponseError struct {
	RespCode string `json:"responseCode"`
	RespMsg  string `json:"responseMessage"`
}

type ResponseList struct {
	ReturnResponseError
	HttpStatusCode string `json:"httpStatusCode"`
}

type ResponseFromDTO struct {
	shared.BaseResponse
	httpStatusCode ResponseList
}

type ListResponseFromJsonFile struct {
	ListResponseFromJsonFile map[string]ResponseList
}

var errorDataList map[string]ResponseList

func ErrorHanlder(ctx *gin.Context) {
	ctx.Next()

	fmt.Println(len(ctx.Errors))
	if len(ctx.Errors) > 0 {
		for _, err := range ctx.Errors {
			resp := ResponseError(err.Error())
			convertHTTPStatusCode, _ := strconv.Atoi(resp.HttpStatusCode)
			if err.Meta != nil {
				resp.RespMsg = fmt.Sprintf(resp.RespMsg, err.Meta)
			}
			ctx.JSON(convertHTTPStatusCode, ReturnResponseError{
				RespCode: resp.RespCode,
				RespMsg:  resp.RespMsg,
			})
			return
		}
	}
}

func ResponseError(respCode string) ResponseList {
	var loadResponse = SearchResponseValueFromJsonFile(respCode)

	return ResponseList{
		ReturnResponseError: ReturnResponseError{
			RespCode: respCode,
			RespMsg:  loadResponse.RespMsg,
		},
		HttpStatusCode: loadResponse.HttpStatusCode,
	}
}

func ResponseSuccess(ctx *gin.Context, respCode string, data interface{}) {

	var loadResponse = SearchResponseValueFromJsonFile(respCode)

	convertRespCode, _ := strconv.Atoi(respCode)
	convertHTTPStatusCode, _ := strconv.Atoi(loadResponse.HttpStatusCode)

	ctx.JSON(convertHTTPStatusCode, ResponseFromDTO{
		BaseResponse: shared.BaseResponse{
			ResponseCode:    convertRespCode,
			ResponseMessage: loadResponse.RespMsg,
			Data:            data,
		},
	})
	return
}

func SearchResponseValueFromJsonFile(resCode string) ResponseList {
	var loadListResponse = errorDataList
	resCodeValue, errResCodeValue := loadListResponse[resCode]
	if errResCodeValue {
		return resCodeValue
	} else {
		return ResponseList{
			ReturnResponseError: ReturnResponseError{
				RespCode: resCode,
				RespMsg:  "Error message is not defined!",
			},
			HttpStatusCode: strconv.Itoa(http.StatusInternalServerError),
		}
	}
}

func LoadErrorListFromJsonFile(pathfilename string) error {
	var file []byte
	var err error

	file, err = os.ReadFile(pathfilename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &errorDataList)
	if err != nil {
		return err
	} else {
		return nil
	}
}
