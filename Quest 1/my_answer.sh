#!/bin/bash
if [ ! -d "the-final-cl-test" ]; then
    git clone https://github.com/01-edu/the-final-cl-test
fi

cd the-final-cl-test || exit
echo "Line 1: Some data" > hidden_info.txt
echo "John Doe" >> hidden_info.txt
echo "Line 3: More data" >> hidden_info.txt
echo "$" >> hidden_info.txt
echo "Line 5: Final data" >> hidden_info.txt
head -n 2 hidden_info.txt | tail -n 1