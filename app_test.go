package main

import "testing"

func TestConvertToInt(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{name: "valid", input: "12", want: 12, wantErr: false},
		{name: "empty", input: "", wantErr: true},
		{name: "non-numeric", input: "abc", wantErr: true},
		{name: "negative", input: "-2", want: -2, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToInt(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ConvertToInt() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Fatalf("ConvertToInt() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestNormalizeVars(t *testing.T) {
	tests := []struct {
		name string
		in   *Vars
		want *Vars
	}{
		{
			name: "all-empty-uses-defaults",
			in:   &Vars{},
			want: &Vars{Name: defaultName, Character: defaultCharacter, Seconds: "0"},
		},
		{
			name: "partial-empty-only-fills-missing",
			in:   &Vars{Name: "alice", Character: "", Seconds: "5"},
			want: &Vars{Name: "alice", Character: defaultCharacter, Seconds: "5"},
		},
		{
			name: "already-set-values-unchanged",
			in:   &Vars{Name: "bob", Character: "neo", Seconds: "9"},
			want: &Vars{Name: "bob", Character: "neo", Seconds: "9"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizeVars(tt.in)
			if *got != *tt.want {
				t.Fatalf("NormalizeVars() = %#v, want %#v", *got, *tt.want)
			}
		})
	}
}

func TestValidateRunTime(t *testing.T) {
	tests := []struct {
		name    string
		sec     int
		wantErr bool
	}{
		{name: "zero", sec: 0, wantErr: false},
		{name: "valid-upper-bound", sec: maxRunTime, wantErr: false},
		{name: "negative", sec: -1, wantErr: true},
		{name: "too-large", sec: maxRunTime + 1, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateRunTime(tt.sec)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ValidateRunTime() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
