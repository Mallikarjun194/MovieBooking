package router

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"lcode/constants"
// 	"lcode/models"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestSetupRouter(t *testing.T) {
// 	tests := []struct {
// 		name       string
// 		statusCode int
// 		body       models.Movie
// 		want       string
// 	}{
// 		{
// 			name:       "Sucessfully create movie entry",
// 			statusCode: http.StatusCreated,
// 			body: models.Movie{
// 				Name:     "Titanic",
// 				Language: "Hindi",
// 				Length:   1,
// 			},
// 		},
// 		{
// 			name:       "Sucessfully create movie entry with body",
// 			statusCode: http.StatusCreated,
// 			body: models.Movie{
// 				Name:     "Titanic",
// 				Language: "Hindi",
// 				Length:   1,
// 			},
// 			want: "id",
// 		},
// 		{
// 			name:       "failes for empty movie name",
// 			statusCode: http.StatusBadRequest,
// 			body: models.Movie{
// 				Language: "Hindi",
// 				Length:   1,
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			router := SetupRouter()

// 			w := httptest.NewRecorder()
// 			jsonValue, _ := json.Marshal(tt.body)
// 			req, _ := http.NewRequest(constants.POST, constants.Movies, bytes.NewBuffer(jsonValue))
// 			router.ServeHTTP(w, req)

// 			if tt.statusCode == http.StatusCreated {
// 				assert.Contains(t, w.Body.String(), tt.want)
// 			}

// 			assert.Equal(t, tt.statusCode, w.Code)

// 		})
// 	}
// }

// func TestGetMovies(t *testing.T) {
// 	tests := []struct {
// 		name       string
// 		statusCode int
// 	}{
// 		{
// 			name:       "Sucessfully get movies list",
// 			statusCode: http.StatusOK,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			router := SetupRouter()

// 			w := httptest.NewRecorder()

// 			req, _ := http.NewRequest(constants.GET, constants.Movies, nil)
// 			router.ServeHTTP(w, req)

// 			assert.Equal(t, tt.statusCode, w.Code)

// 		})
// 	}
// }

// func TestUpdateRouter(t *testing.T) {
// 	tests := []struct {
// 		name       string
// 		statusCode int
// 		body       models.Movie
// 		want       string
// 	}{
// 		{
// 			name:       "Sucessfully create movie entry",
// 			statusCode: http.StatusOK,
// 			body: models.Movie{
// 				Name:     "Titanic",
// 				Language: "Hindi",
// 				Length:   1,
// 				Id:       "test",
// 			},
// 		},

// 		{
// 			name:       "failes for empty movie name",
// 			statusCode: http.StatusBadRequest,
// 			body: models.Movie{
// 				Language: "Hindi",
// 				Length:   1,
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			router := SetupRouter()

// 			w := httptest.NewRecorder()
// 			jsonValue, _ := json.Marshal(tt.body)
// 			req, _ := http.NewRequest(constants.PATCH, constants.Movies, bytes.NewBuffer(jsonValue))
// 			router.ServeHTTP(w, req)

// 			if tt.statusCode == http.StatusCreated {
// 				assert.Contains(t, w.Body.String(), tt.want)
// 			}

// 			fmt.Printf("w.Body.String(): %v\n", w.Body.String())
// 			assert.Equal(t, tt.statusCode, w.Code)

// 		})
// 	}
// }

// // func TestDeleteRouter(t *testing.T) {
// // 	tests := []struct {
// // 		name       string
// // 		statusCode int
// // 		body       models.Movie
// // 		want       string
// // 	}{
// // 		{
// // 			name:       "Sucessfully create movie entry",
// // 			statusCode: http.StatusOK,
// // 			body: models.Movie{
// // 				Name:     "Titanic",
// // 				Language: "Hindi",
// // 				Length:   1,
// // 				Id:       "test",
// // 			},
// // 		},

// // 		{
// // 			name:       "failes for empty movie name",
// // 			statusCode: http.StatusBadRequest,
// // 			body: models.Movie{
// // 				Language: "Hindi",
// // 				Length:   1,
// // 			},
// // 		},
// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			router := SetupRouter()

// // 			w := httptest.NewRecorder()
// // 			jsonValue, _ := json.Marshal(tt.body)
// // 			req, _ := http.NewRequest(constants.PATCH, constants.Movies, bytes.NewBuffer(jsonValue))
// // 			router.ServeHTTP(w, req)

// // 			if tt.statusCode == http.StatusCreated {
// // 				assert.Contains(t, w.Body.String(), tt.want)
// // 			}

// // 			fmt.Printf("w.Body.String(): %v\n", w.Body.String())
// // 			assert.Equal(t, tt.statusCode, w.Code)

// // 		})
// // 	}
// // }
