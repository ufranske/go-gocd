package gocd

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestApiVersion_Int(t *testing.T) {
	for _, tt := range []struct {
		name string
		av   ApiVersion
		want int
	}{
		{
			name: "v1",
			av:   ApiVersion0,
			want: 0,
		},
		{
			name: "v5",
			av:   ApiVersion5,
			want: 5,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.av.Int()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestApiVersion_LessThan(t *testing.T) {
	type args struct {
		version ApiVersion
	}
	for _, tt := range []struct {
		name string
		av   ApiVersion
		args args
		want bool
	}{
		{
			name: "success",
			av:   ApiVersion0,
			args: args{
				version: ApiVersion1,
			},
			want: true,
		},
		{
			name: "fail",
			av:   ApiVersion1,
			args: args{
				version: ApiVersion0,
			},
			want: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.av.LessThan(tt.args.version)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestApiVersion_Equal(t *testing.T) {
	type args struct {
		version ApiVersion
	}
	for _, tt := range []struct {
		name string
		av   ApiVersion
		args args
		want bool
	}{
		{
			name: "equal",
			av:   ApiVersion1,
			args: args{
				version: ApiVersion1,
			},
			want: true,
		},
		{
			name: "not-equal",
			av:   ApiVersion1,
			args: args{
				version: ApiVersion1,
			},
			want: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.av.Equal(tt.args.version)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestApiVersion_String(t *testing.T) {
	for _, tt := range []struct {
		name string
		av   ApiVersion
		want string
	}{
		{
			name: "v1",
			av:   ApiVersion1,
			want: "application/vnd.go.cd.v1+json",
		},
		{
			name: "v4",
			av:   ApiVersion4,
			want: "application/vnd.go.cd.v4+json",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.av.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewApiVersionRange(t *testing.T) {
	type args struct {
		min ApiVersion
		max ApiVersion
	}
	for _, tt := range []struct {
		name string
		args args
		want *ApiVersionRange
	}{
		{
			name: "basic",
			args: args{
				min: ApiVersion0,
				max: ApiVersion1,
			},
			want: &ApiVersionRange{
				min: ApiVersion0,
				max: ApiVersion1,
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := NewApiVersionRange(tt.args.min, tt.args.max)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestApiVersionRange_Contains(t *testing.T) {
	type fields struct {
		min ApiVersion
		max ApiVersion
	}
	type args struct {
		version ApiVersion
	}
	for _, tt := range []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "success",
			fields: fields{
				min: ApiVersion0,
				max: ApiVersion3,
			},
			args: args{
				version: ApiVersion2,
			},
			want: true,
		},
		{
			name: "fail",
			fields: fields{
				min: ApiVersion0,
				max: ApiVersion3,
			},
			args: args{
				version: ApiVersion5,
			},
			want: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			avr := &ApiVersionRange{
				min: tt.fields.min,
				max: tt.fields.max,
			}
			got := avr.Contains(tt.args.version)
			assert.Equal(t, tt.want, got)
		})
	}
}
