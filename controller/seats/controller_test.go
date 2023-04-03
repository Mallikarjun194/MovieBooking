package seats

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

func TestSeatController_AddSeat(t *testing.T) {
	tests := []struct {
		name string
		body models.Seat
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Seat{
			Type:      "A",
			Number:    10,
			TheaterID: "theater-1",
		},
		mock: func() repository.RepositoryI {
			ri := mocks.RepositoryI{}
			ri.On("Create", mock.Anything).Return(nil)
			return &ri
		},
		want: http.StatusCreated,
	},
		{
			name: "failes for empty theaterID",
			want: http.StatusBadRequest,
			mock: func() repository.RepositoryI {
				ri := mocks.RepositoryI{}
				ri.On("Create", mock.Anything).Return(nil)
				return &ri
			},
			body: models.Seat{
				Type:   "A",
				Number: 10,
			},
		},
	}
	router := gin.Default()
	m := &SeatController{}
	router.GET(rPath, m.AddSeat)
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

func TestSeatController_GetSeats(t *testing.T) {
	tests := []struct {
		name string
		body models.Seat
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Seat{ID: "test"},
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
			m := &SeatController{Repository: tt.mock()}
			router.GET(rPath, m.GetSeats)
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

func TestSeatController_DeleteSeat(t *testing.T) {

	tests := []struct {
		name string
		body models.Seat
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Seat{
			Type:      "A",
			Number:    10,
			TheaterID: "theater-1",
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
			m := &SeatController{Repository: tt.mock()}
			router.GET(rPath, m.DeleteSeat)
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
