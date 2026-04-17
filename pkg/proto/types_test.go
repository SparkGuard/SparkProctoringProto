package proto

import "testing"

func TestFormatStreamType(t *testing.T) {
	cases := []struct {
		name    string
		webcam  bool
		channel uint8
		want    string
	}{
		{"primary screen", false, 0, "screen"},
		{"second screen", false, 1, "screen_1"},
		{"third screen", false, 2, "screen_2"},
		{"primary webcam", true, 0, "webcam"},
		{"fourth webcam", true, 3, "webcam_3"},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := FormatStreamType(tc.webcam, tc.channel); got != tc.want {
				t.Errorf("FormatStreamType(%v, %d) = %q, want %q",
					tc.webcam, tc.channel, got, tc.want)
			}
		})
	}
}
