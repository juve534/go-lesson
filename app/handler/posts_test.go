package handler_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/juve534/go-lesson/app/handler"
	"github.com/juve534/go-lesson/app/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mock struct {
	t          *testing.T
	GetFunc    func(postID string) (*models.Posts, error)
	CreateFunc func(posts *models.Posts) error
}

func (m *mock) GetPostById(postID string) (*models.Posts, error) {
	return m.GetFunc(postID)
}

func (m *mock) CreatePost(posts *models.Posts) error {
	return m.CreateFunc(posts)
}

func TestPostHandler_PostIndex(t *testing.T) {
	type args struct {
		req     *http.Request
		GetFunc func(postID string) (*models.Posts, error)
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			"パスパラメータの値に紐づくデータが存在するとき、データが返却される",
			args{
				req: func() *http.Request {
					return httptest.NewRequest("POST", "http://target/1", nil)
				}(),
				GetFunc: func(postID string) (*models.Posts, error) {
					return &models.Posts{
						ID:    1,
						Title: "hoge-Title",
						Body:  "hoge-Body",
					}, nil
				},
			},
			`{"id":1,"title":"hoge-Title","body":"hoge-Body"}
`,
			nil,
		},
		{
			"パスパラメータの値に紐づくデータが存在しないとき、空のレスポンスが返却される",
			args{
				req: func() *http.Request {
					return httptest.NewRequest("POST", "http://target/1", nil)
				}(),
				GetFunc: func(postID string) (*models.Posts, error) {
					return nil, nil
				},
			},
			`null
`,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mock{
				t:       t,
				GetFunc: tt.args.GetFunc,
			}
			h := handler.NewPostHandler(mock)
			rw := httptest.NewRecorder()
			h.PostIndex(rw, tt.args.req)
			gotBody, _ := ioutil.ReadAll(rw.Body)
			if diff := cmp.Diff(string(gotBody), tt.want); diff != "" {
				t.Errorf("PostIndex() Body\n%s", diff)
			}
		})
	}
}
