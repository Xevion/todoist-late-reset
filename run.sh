set -ex

go build -o out ./cmd/${1:-main}/
chmod +x ./out
./out
rm -f ./out