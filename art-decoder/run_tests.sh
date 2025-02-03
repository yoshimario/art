#!/bin/bash

# Define the path to your compiled executable
executable="./art-decoder"

# Define paths for the encoded and expected decoded files
encoded_dir="files/encode"
decoded_dir="files/decode"

# Check if the executable exists
if [ ! -x "$executable" ]; then
  echo "Error: Executable $executable not found or not executable."
  exit 1
fi

echo "🔍 Running inline tests..."

# Inline test cases
declare -A tests=(
  ["[5 #][5 -_]-[5 #]"]="#####-_-_-_-_-_-#####"
  ["[3 @][2 !]"]="@@@!!"
  ["[5 #][5 -_]-[5 #]]"]="Error: Extra closing bracket found"
  ["[5 #]5 -_]-[5 #]"]="Error: Missing opening bracket"
  ["[5#][5 -_]-[5 #]"]="Error: Invalid format inside brackets (expected '[count char]')"
  ["5 #[5 -_]-5 #"]="Error: Missing opening bracket"
)

# Run inline tests
for input in "${!tests[@]}"; do
  expected="${tests[$input]}"
  output=$("$executable" -ml <<< "$input")  # Run the decoder with inline input

  if [[ "$output" == "$expected" ]]; then
    echo "✅ Inline Test passed: $input"
  else
    echo "❌ Inline Test failed: $input"
    echo "Expected: $expected"
    echo "Got: $output"
    exit 1
  fi
done

echo "✅ All inline tests passed!"
echo ""
echo "🔍 Running file-based tests..."

# Array of test case file names (without extensions)
test_files=("cats" "kood" "lion" "plane")

# Loop over each test case (file-based)
for name in "${test_files[@]}"; do
  encoded_file="${encoded_dir}/${name}.encoded.txt"
  expected_file="${decoded_dir}/${name}.art.txt"

  # Ensure files exist
  if [ ! -f "$encoded_file" ]; then
    echo "❌ Error: Encoded file $encoded_file not found."
    continue
  fi
  if [ ! -f "$expected_file" ]; then
    echo "❌ Error: Expected decoded file $expected_file not found."
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
    echo "Expected output from $expected_file:"
    echo "$expected_output"
    echo "Got:"
    echo "$output"
    exit 1
  fi
done

echo "🎉 All tests (inline + file-based) completed successfully!"