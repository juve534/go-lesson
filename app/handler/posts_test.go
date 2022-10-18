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

func NewMock() models.PostsModel {
	return &mock{}
}

type mock struct {
}

func (m *mock) GetPostById(postID string) (*models.Posts, error) {
	return &models.Posts{
		ID:    1,
		Title: "hoge-Title",
		Body:  "hoge-Body",
	}, nil
}

func (m mock) CreatePost(posts *models.Posts) error {
	posts.ID = 1
	return nil
}

func TestPostHandler_PostIndex(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			"パスパラメータの値が正しいとき",
			args{
				func() *http.Request {
					return httptest.NewRequest("POST", "http://target/1", nil)
				}(),
			},
			`{"id":1,"title":"hoge-Title","body":"hoge-Body"}
`,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler.NewPostHandler(NewMock())
			rw := httptest.NewRecorder()
			h.PostIndex(rw, tt.args.req)
			gotBody, _ := ioutil.ReadAll(rw.Body)
			if diff := cmp.Diff(string(gotBody), tt.want); diff != "" {
				t.Errorf("PostIndex() Body\n%s", diff)
			}
		})
	}
}
