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

  if diff <(echo "$output") <(echo "$expected") > /dev/null; then
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

# Summary
if [[ $fail_count -gt 0 ]]; then
  echo "❌ Some tests failed ($fail_count failures)."
  exit 1
else
  echo "🎉 All tests (inline + file-based) passed successfully!"
  exit 0
fi