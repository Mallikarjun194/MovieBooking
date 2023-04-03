package theater

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

func TestTheaterController_AddTheater(t *testing.T) {
	tests := []struct {
		name string
		body models.Theater
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Theater{},
		mock: func() repository.RepositoryI {
			ri := mocks.RepositoryI{}
			ri.On("Create", mock.Anything).Return(nil)
			return &ri
		},
		want: http.StatusCreated,
	},
		{
			name: "failes for empty theater name",
			want: http.StatusBadRequest,
			mock: func() repository.RepositoryI {
				ri := mocks.RepositoryI{}
				ri.On("Create", mock.Anything).Return(nil)
				return &ri
			},
			body: models.Theater{},
		},
	}
	router := gin.Default()
	m := &TheaterController{}
	router.GET(rPath, m.AddTheater)
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

func Testtheatercontroller_Gettheaters(t *testing.T) {
	tests := []struct {
		name string
		body models.Theater
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Theater{ID: "test"},
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
			m := &TheaterController{Repository: tt.mock()}
			router.GET(rPath, m.GetTheaters)
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

func Testtheatercontroller_Updatetheater(t *testing.T) {
	tests := []struct {
		name string
		body models.Theater
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Theater{
			ID:   "id",
			Name: "Titanic",
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
			m := &TheaterController{Repository: tt.mock()}
			router.GET(rPath, m.UpdateTheater)
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

func Testtheatercontroller_Deletetheater(t *testing.T) {

	tests := []struct {
		name string
		body models.Theater
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Theater{
			ID:   "id",
			Name: "Titanic",
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
			m := &TheaterController{Repository: tt.mock()}
			router.GET(rPath, m.DeleteTheater)
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
