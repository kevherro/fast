package math

import (
	stdMath "math"
	"testing"
)

func TestQRsqrt(t *testing.T) {
	var tests = []struct {
		name      string
		input     float32
		want      float32
		wantInf   bool
		wantNaN   bool
		wantPanic bool
	}{
		{
			name:  "Inverse square root of 4",
			input: 4,
			want:  0.5,
		},
		{
			name:  "Inverse square root of 1",
			input: 1,
			want:  1,
		},
		{
			name:  "Inverse square root of 0.25",
			input: 0.25,
			want:  2,
		},
		{
			name:  "Inverse square root of 3",
			input: 3,
			want:  0.577,
		},
		{
			name:  "Inverse square root of 5",
			input: 5,
			want:  0.447,
		},
		{
			name:    "Inverse square root of 0.0",
			input:   0.0,
			wantInf: true,
		},
		{
			name:    "Inverse square root of -1",
			input:   -1,
			wantInf: true,
		},
		{
			name:  "Inverse square root of a very large number",
			input: 1.0e30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !tt.wantPanic {
					t.Errorf("QRsqrt() panicked, but didn't want it to")
				}
			}()

			got := QRsqrt(tt.input)

			if stdMath.IsInf(float64(got), 1) != tt.wantInf {
				t.Errorf("QRsqrt() = %v, want %v", got, stdMath.Inf(1))
			}

			if stdMath.IsNaN(float64(got)) != tt.wantNaN {
				t.Errorf("QRsqrt() = %v, want NaN", got)
			}

			const epsilon = 1.0e-3
			if !tt.wantInf && !tt.wantNaN && stdMath.Abs(float64(got-tt.want)) > epsilon {
				t.Errorf("QRsqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkQRsqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QRsqrt(4.0)
	}
}
