protoc -I=schemas --go_out=node schemas/fc.proto


for file in abi/*.json; do
  filename=$(basename "$file" .json)
  abigen --abi "$file" --pkg contracts --type "$filename" --out "node/contracts/$filename.go"
done




