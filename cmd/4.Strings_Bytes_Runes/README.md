# Strings, Runes & Bytes in Go

How Go represents and manipulates strings. The interplay between strings, their underlying byte arrays, and runes (Unicode code points).

## 📜 Discussed

- **String Representation:**  
  Understanding that Go strings are stored as an array of bytes (UTF-8 encoded).  
  Indexing a string returns a byte (uint8), which may not represent the full character when dealing with non-ASCII symbols.

- **UTF-8 & Non-ASCII Characters:**  
  Explore how multi-byte characters (like the acute 'é' in "résumé") are encoded, and why iterating with the `range` keyword correctly decodes them.

- **Runes:**  
  Learn how converting a string to a slice of runes gives you proper Unicode code points (int32), making indexing and iteration more intuitive for non-ASCII text.

- **String Concatenation:**  
  See the difference between naive string concatenation (which creates new strings on every append) and efficient concatenation using Go's `strings.Builder`.

## How to Run

   ```bash
   go run cmd/4.Strings_Bytes_Runes/main.go
