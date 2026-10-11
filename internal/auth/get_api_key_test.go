package auth

// func TestSplit(t *testing.T) {
// 	t.Run("returns error when authorization header is missing", func(t *testing.T) {
// 		headers := http.Header{}
// 		_, err := auth.GetAPIKey(headers)
// 		if err == nil {
// 			t.Fatalf("expected error, got nil")
// 		}
// 		if err != auth.ErrNoAuthHeaderIncluded {
// 			t.Fatalf("expected error %v, got %v", auth.ErrNoAuthHeaderIncluded, err)
// 		}
// 	})

// 	t.Run("returns error when authorization header is malformed", func(t *testing.T) {
// 		headers := http.Header{}
// 		headers.Set("Authorization", "Bearer my-api-key")
// 		_, err := auth.GetAPIKey(headers)
// 		if err == nil {
// 			t.Fatalf("expected error, got nil")
// 		}
// 		if err.Error() != "malformed authorization header" {
// 			t.Fatalf("expected error 'malformed authorization header', got %v", err)
// 		}
// 	})
// 	t.Run("returns error when authorization header is empty", func(t *testing.T) {
// 		headers := http.Header{}
// 		headers.Set("Authorization", "")
// 		_, err := auth.GetAPIKey(headers)
// 		if err == nil {
// 			t.Fatalf("expected error, got nil")
// 		}
// 		if err != auth.ErrNoAuthHeaderIncluded {
// 			t.Fatalf("expected error %v, got %v", auth.ErrNoAuthHeaderIncluded, err)
// 		}
// 	})
// }
