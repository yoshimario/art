#!/bin/bash

# Array of test case base names
test_files=("cats" "kood" "lion" "plane")

# Define paths for the encoded and expected decoded files
encoded_dir="files/encode"
decoded_dir="files/decode"
executable="./art-decoder"  # Adjust this if your binary has a different name

# Check if the executable exists
if [ ! -x "$executable" ]; then
  echo "Error: Executable $executable not found or not executable."
  exit 1
fi

# Loop over each test case
for name in "${test_files[@]}"; do
  encoded_file="${encoded_dir}/${name}.encoded.txt"
  expected_file="${decoded_dir}/${name}.art.txt"
  
  # Verify that files exist
  if [ ! -f "$encoded_file" ]; then
    echo "Encoded file $encoded_file not found!"
    continue
  fi
  
  if [ ! -f "$expected_file" ]; then
    echo "Expected file $expected_file not found!"
    continue
  fi
  
  # Run the program using input redirection
  output=$("$executable" -ml < "$encoded_file")
  
  # Compare the output with the expected output using diff
  diff_output=$(diff <(echo "$output") "$expected_file")
  if [ $? -eq 0 ]; then
    echo "✅ Test passed for $name"
  else
    echo "❌ Test failed for $name"
    echo "Differences:"
    echo "$diff_output"
  fi
done
