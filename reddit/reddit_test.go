package reddit

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPicture(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "no picture",
			args: args{
				url: "https://test.com/hawdhawhdh",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "is picture",
			args: args{
				url: "https://test.com/hawdhawhdh.png",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uri, err := url.Parse(tt.args.url)
			assert.NoError(t, err)
			got, err := isPicture(uri)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
