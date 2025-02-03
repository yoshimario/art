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
  "[5#][5 -_]-[5 #]"
  "5 #[5 -_]-5 #"
)

expected_outputs=(
  "#####-_-_-_-_-_-#####"
  "@@@!!"
  "Error: Extra closing bracket found"
  "Error: Missing opening bracket"
  "Error: Invalid format inside brackets (expected '[count char]')"
  "Error: Missing opening bracket"
)

# Track test failures
fail_count=0

# Run inline tests
for i in "${!test_inputs[@]}"; do
  input="${test_inputs[$i]}"
  expected="${expected_outputs[$i]}"
  output=$("$executable" -ml <<< "$input")  # Run the decoder with inline input

  if [[ "$output" == "$expected" ]]; then
    echo "✅ Inline Test passed: $input"
  else
    echo "❌ Inline Test failed: $input"
    echo "   Expected: $expected"
    echo "   Got:      $output"
    fail_count=$((fail_count + 1))
  fi
done

echo "✅ All inline tests completed!"
echo ""
echo "🔍 Running file-based tests..."

# Array of test case file names (without extensions)
test_files=("cats" "kood" "lion" "plane")

# Run file-based tests
for name in "${test_files[@]}"; do
  encoded_file="${encoded_dir}/${name}.encoded.txt"
  expected_file="${decoded_dir}/${name}.art.txt"

  # Ensure files exist
  if [ ! -f "$encoded_file" ]; then
    echo "❌ Error: Encoded file $encoded_file not found."
    fail_count=$((fail_count + 1))
    continue
  fi
  if [ ! -f "$expected_file" ]; then
    echo "❌ Error: Expected decoded file $expected_file not found."
    fail_count=$((fail_count + 1))
    continue
  fi

  # Run the decoder with -ml for multi-line decoding
  output=$("$executable" -ml < "$encoded_file")

  # Read the expected output from the file
  expected_output=$(cat "$expected_file")

  # Compare output
  if [[ "$output" == "$expected_output" ]]; then
    echo "✅ File Test passed: $name"
  else
    echo "❌ File Test failed: $name"
    echo "   Expected output from $expected_file:"
    echo "$expected_output"
    echo "   Got:"
    echo "$output"
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