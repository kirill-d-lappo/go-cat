echo "test message from file 1 - #1
test message from file 1 - #2
" > in1

echo "test message from file 2 - Solve The Mystery" > in2

go build && ./cat in1 in2 > out.copy

cat in1 in2 > out.original

vimdiff out.original out.copy