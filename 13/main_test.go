package main

// func TestIsList(t *testing.T) {
// 	tests := []struct {
// 		input string
// 		want  bool
// 	}{
// 		{"[]", true},
// 		{"[1]", true},
// 		{"[2,3,4]", true},
// 		{"1", false},
// 	}
// 	for _, tt := range tests {
// 		name := tt.input
// 		t.Run(name, func(t *testing.T) {
// 			ans := isList(tt.input)
// 			if ans != tt.want {
// 				t.Errorf("got %v, want %v", ans, tt.want)
// 			}
// 		})
// 	}
// }

// func listEqual(a, b []string) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}
// 	for i := 0; i < len(a); i++ {
// 		if a[i] != b[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func TestToList(t *testing.T) {
// 	tests := []struct {
// 		input string
// 		want  []string
// 	}{
// 		{"[]", []string{}},
// 		{"[1]", []string{"1"}},
// 		{"[2,3,4]", []string{"2", "3", "4"}},
// 	}
// 	for _, tt := range tests {
// 		name := tt.input
// 		t.Run(name, func(t *testing.T) {
// 			ans := toList(tt.input)
// 			if !listEqual(ans, tt.want) {
// 				t.Errorf("got %v, want %v", ans, tt.want)
// 			}
// 		})
// 	}

// }

// func TestIsInt(t *testing.T) {
// 	tests := []struct {
// 		input string
// 		want  bool
// 	}{
// 		{"1", true},
// 		{"[]", false},
// 		{"[1]", false},
// 		{"[2,3,4]", false},
// 	}
// 	for _, tt := range tests {
// 		name := tt.input
// 		t.Run(name, func(t *testing.T) {
// 			ans := isInt(tt.input)
// 			if ans != tt.want {
// 				t.Errorf("got %v, want %v", ans, tt.want)
// 			}
// 		})
// 	}
// }

// func TestParseInt(t *testing.T) {
// 	tests := []struct {
// 		input   string
// 		integer int
// 		offset  int
// 	}{
// 		{"1", 1, 1},
// 		{"12", 12, 2},
// 		{"1,2,3]", 1, 1},
// 		{"123]", 123, 3},
// 	}
// 	for _, tt := range tests {
// 		name := tt.input
// 		t.Run(name, func(t *testing.T) {
// 			i, o := parseInt(tt.input)
// 			if i != tt.integer || o != tt.offset {
// 				t.Errorf("got (%v,%v), want (%v,%v)", i, o, tt.integer, tt.offset)
// 			}
// 		})
// 	}
// }
