package error

import (
	"errors"
	"net/http"
	"time"

	"site/api/internal/logger"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(logger logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err != nil {
			fieldsWithError, _ := c.Get(FieldsWithErrorKey)
			castFields := make([]ErrorField, 0)
			if !(fieldsWithError == nil) {
				castFields = fieldsWithError.([]ErrorField)
			}
			createJSONMessage(c, logger, err.Err, castFields)
			return
		}
	}
}

func createJSONMessage(ctx *gin.Context, log logger.Logger, err error, fieldsWithError []ErrorField) {
	errorMessage := ErrorMessage{}
	errorMessage.Message = err.Error()
	errorMessage.Timestamp = time.Now().Format(DateLayout)
	errorMessage.Path = ctx.FullPath()
	errorMessage.Fields = fieldsWithError
	statusCode := resolveStatusCode(err)

	showLog(errorMessage.Message, statusCode, log)
	ctx.JSON(
		statusCode,
		errorMessage,
	)
}

func resolveStatusCode(err error) int {
	for code, errGroup := range statusCodes() {
		for _, possibleError := range errGroup {
			if errors.Is(err, possibleError) {
				return code
			}
		}
	}
	return http.StatusInternalServerError
}

func showLog(message string, statusCode int, log logger.Logger) {
	if statusCode == http.StatusInternalServerError {
		log.Error(message)
		return
	}
	log.Warning(message)
}
