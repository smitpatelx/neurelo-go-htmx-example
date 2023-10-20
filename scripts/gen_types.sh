# Generate types
spec_file="./specs/booking-system-859e7e1cf8e5-rest.json"
output_dir="./src/oapi.gen.go"
config_file="./openapi-config.json"

rm -rf $output_dir
# openapi-generator generate -i $spec_file -g go -o $output_dir -c $config_file
oapi-codegen -package oapi $spec_file > $output_dir