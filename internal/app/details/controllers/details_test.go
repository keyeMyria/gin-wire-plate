package controllers

import (
	"gin-wire-plate/internal/pkg/models"
	"gin-wire-plate/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

var r *gin.Engine
var configFile = flag.String("f", "details.yml", "set config file which viper will loading.")

func setup() {
	r = gin.New()
}

func TestDetailsController_Get(t *testing.T) {
	flag.Parse()
	setup()

	sto := new(mocks.DetailsRepository)

	sto.On("Get", mock.AnythingOfType("uint64")).Return(func(ID uint64) (p *models.Detail) {
		return &models.Detail{
			ID: ID,
		}
	}, func(ID uint64) error {
		return nil
	})

	c, err := CreateDetailsController(*configFile, sto)
	if err != nil {
		t.Fatalf("create product serviceerror,%+v", err)
	}

	r.GET("/proto/:id", c.Get)

	tests := []struct {
		name     string
		id       uint64
		expected uint64
	}{
		{"1", 1, 1},
		{"2", 2, 2},
		{"3", 3, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uri := fmt.Sprintf("/proto/%d", test.id)
			// 构造get请求
			req := httptest.NewRequest("GET", uri, nil)
			// 初始化响应
			w := httptest.NewRecorder()

			// 调用相应的controller接口
			r.ServeHTTP(w, req)

			// 提取响应
			rs := w.Result()
			defer func() {
				_ = rs.Body.Close()
			}()

			// 读取响应body
			body, _ := ioutil.ReadAll(rs.Body)
			p := new(models.Detail)
			err := json.Unmarshal(body, p)
			if err != nil {
				t.Errorf("unmarshal response body error:%v", err)
			}

			assert.Equal(t, test.expected, p.ID)
		})
	}

}
