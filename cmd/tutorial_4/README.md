# Strings, Runes & Bytes in Go

How Go represents and manipulates strings. The interplay between strings, their underlying byte arrays, and runes (Unicode code points).

## What You'll Learn

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

1. **Clone or copy** this code into your project.
2. **Navigate** to the directory containing `main.go`.
3. **Run the program:**

   ```bash
   go run main.go
