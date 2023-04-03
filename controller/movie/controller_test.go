package movie

import (
	"MovieBooking/mocks"
	"MovieBooking/models"
	"MovieBooking/repository"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var rPath = "/test"

func TestMovieController_AddMovie(t *testing.T) {
	tests := []struct {
		name string
		body models.Movie
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Movie{
			Name:     "Titanic",
			Language: "Hindi",
			Length:   1,
		},
		mock: func() repository.RepositoryI {
			ri := mocks.RepositoryI{}
			ri.On("Create", mock.Anything).Return(nil)
			return &ri
		},
		want: http.StatusCreated,
	},
		{
			name: "failes for empty movie name",
			want: http.StatusBadRequest,
			mock: func() repository.RepositoryI {
				ri := mocks.RepositoryI{}
				ri.On("Create", mock.Anything).Return(nil)
				return &ri
			},
			body: models.Movie{
				Language: "Hindi",
				Length:   2,
			},
		},
	}
	router := gin.Default()
	m := &MovieController{}
	router.GET(rPath, m.AddMovie)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.Repository = tt.mock()
			b, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("GET", rPath, bytes.NewReader(b))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			t.Logf("status: %d", w.Code)
			t.Logf("response: %s", w.Body.String())
			assert.Equal(t, tt.want, w.Code)
		})
	}
}

func TestMovieController_GetMovies(t *testing.T) {
	tests := []struct {
		name string
		body models.Movie
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Movie{ID: "test"},
		mock: func() repository.RepositoryI {
			ri := mocks.RepositoryI{}
			ri.On("QueryAll", mock.Anything).Return(nil)
			return &ri
		},
		want: http.StatusOK,
	}}

	router := gin.Default()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MovieController{Repository: tt.mock()}
			router.GET(rPath, m.GetMovies)
			b, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("GET", rPath, bytes.NewReader(b))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			t.Logf("status: %d", w.Code)
			t.Logf("response: %s", w.Body.String())
			assert.Equal(t, tt.want, w.Code)
		})
	}
}

func TestMovieController_UpdateMovie(t *testing.T) {
	tests := []struct {
		name string
		body models.Movie
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Movie{
			ID:       "id",
			Name:     "Titanic",
			Language: "Hindi",
			Length:   2,
		},
		mock: func() repository.RepositoryI {
			ri := mocks.RepositoryI{}
			ri.On("Update", mock.Anything).Return(nil)
			return &ri
		},
		want: http.StatusOK,
	}}

	router := gin.Default()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MovieController{Repository: tt.mock()}
			router.GET(rPath, m.UpdateMovie)
			b, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("GET", rPath, bytes.NewReader(b))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			t.Logf("status: %d", w.Code)
			t.Logf("response: %s", w.Body.String())
			assert.Equal(t, tt.want, w.Code)
		})
	}
}

func TestMovieController_DeleteMovie(t *testing.T) {

	tests := []struct {
		name string
		body models.Movie
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Movie{
			ID:       "id",
			Name:     "Titanic",
			Language: "Hindi",
			Length:   2,
		},
		mock: func() repository.RepositoryI {
			ri := mocks.RepositoryI{}
			ri.On("Delete", mock.Anything).Return(nil)
			return &ri
		},
		want: http.StatusOK,
	}}

	router := gin.Default()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MovieController{Repository: tt.mock()}
			router.GET(rPath, m.DeleteMovie)
			b, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("GET", rPath, bytes.NewReader(b))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			t.Logf("status: %d", w.Code)
			t.Logf("response: %s", w.Body.String())
			assert.Equal(t, tt.want, w.Code)
		})
	}
}
