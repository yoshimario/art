#!/bin/bash

# Define the path to your compiled executable
executable="./art-decoder"

# Define paths for the encoded and expected decoded files
encoded_dir="files/encode"
decoded_dir="files/decode"

# Check if the executable exists
if [ ! -x "$executable" ]; then
  echo "❌ Error: Executable $executable not found or not executable."
  exit 1
fi

echo "🔍 Running inline tests..."

# Define test cases using two separate arrays
test_inputs=(
  "[5 #][5 -_]-[5 #]"
  "[3 @][2 !]"
  "[5 #][5 -_]-[5 #]]"
  "[5 #]5 -_]-[5 #]"
  "[5 #][5 -_-[5 #]"
  "[5#][5 -_]-[5 #]"
  "5 #[5 -_]-5 #"
  "5 #[5 -_]-5 #"
)

expected_outputs=(
  "#####-_-_-_-_-_-#####"
  "@@@!!"
  "Error: Extra closing bracket found"
  "Error: Missing opening bracket"
  "Error: Missing closing bracket"
  "Error: Invalid format inside brackets (expected '[count char]')"
  "Error: Missing opening bracket"
  "Error: Missing opening bracket"
)

# Track test failures
fail_count=0

# Run inline tests
for i in "${!test_inputs[@]}"; do
  input="${test_inputs[$i]}"
  expected="${expected_outputs[$i]}"

  # Run the decoder with multi-line flag and capture both stdout and stderr
  output_with_color=$("$executable" -ml <<< "$input" 2>&1)
  
  # Create a copy without color codes for comparison only
  output_for_comparison=$(echo "$output_with_color" | sed 's/\x1b\[[0-9;]*m//g')

  # Remove trailing newlines for comparison (cross-platform solution)
  output_for_comparison=$(echo -n "$output_for_comparison" | perl -pe 'chomp if eof')
  expected=$(echo -n "$expected" | perl -pe 'chomp if eof')

  if [ "$output_for_comparison" == "$expected" ]; then
    echo -e "✅ Inline Test passed: $input"
    # Display the colored output if test passed
    echo -e "   Result: $output_with_color"
  else
    echo -e "❌ Inline Test failed: $input"
    echo -e "   Expected: $expected"
    echo -e "   Got (without colors): $output_for_comparison"
    echo -e "   Got (with colors): $output_with_color"
    fail_count=$((fail_count + 1))
  fi
done

echo "✅ All inline tests completed!"
echo ""
echo "🔍 Running file-based tests..."

# Check if the encoded and decoded directories exist
if [ ! -d "$encoded_dir" ] || [ ! -d "$decoded_dir" ]; then
  echo "❌ Error: Directories $encoded_dir or $decoded_dir not found."
  exit 1
fi

# Run file-based tests
for encoded_file in "$encoded_dir"/*.encoded.txt; do
  filename=$(basename "$encoded_file")
  base_name="${filename%.encoded.txt}"  # Remove .encoded.txt suffix
  decoded_file="${decoded_dir}/${base_name}.art.txt"  # Match expected .art.txt file

  if [ ! -f "$decoded_file" ]; then
    echo "❌ Error: Expected decoded file $decoded_file not found."
    fail_count=$((fail_count + 1))
    continue
  fi

  # Run the decoder on the encoded file and capture output with colors
  output_with_color=$("$executable" -ml < "$encoded_file" 2>&1)
  
  # Create a copy without color codes for comparison only
  output_for_comparison=$(echo "$output_with_color" | sed 's/\x1b\[[0-9;]*m//g')

  # Read the expected decoded content
  expected=$(cat "$decoded_file")

  # Remove trailing newlines for comparison
  output_for_comparison=$(echo -n "$output_for_comparison" | perl -pe 'chomp if eof')
  expected=$(echo -n "$expected" | perl -pe 'chomp if eof')

  if [ "$output_for_comparison" == "$expected" ]; then
    echo -e "✅ File Test passed: $filename"
    # Display the colored output if test passed
    echo -e "   Result: $output_with_color"
  else
    echo -e "❌ File Test failed: $filename"
    echo -e "   Expected: $expected"
    echo -e "   Got (without colors): $output_for_comparison"
    echo -e "   Got (with colors): $output_with_color"
    fail_count=$((fail_count + 1))
  fi
done

# Summary
if [[ $fail_count -gt 0 ]]; then
  echo "❌ Some tests failed ($fail_count failures)."
  exit 1
else
  echo "🎉 All tests (inline + file-based) passed successfully!"
  exit 0
fi