package tests

import (
	"example/weather-rest-api/config"
	"example/weather-rest-api/server"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// TestSystemHealth tests the system health API
func TestSystemHealth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.Init("test")

	router := server.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

// TestToWeatherForecastFailure tests the failure of the weather forecast API
func TestNotSubscribeToWeatherForecast(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.Init("test")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "/api/v1/weather", httpmock.NewJsonResponderOrPanic(403, httpmock.File("../forbidden.json")))
}

// TestToWeatherForecastSuccess tests the success of the weather forecast API
func TestToWeatherForecastSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.Init("test")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "/api/v1/weather", httpmock.NewJsonResponderOrPanic(200, httpmock.File("../response.json")))

}
