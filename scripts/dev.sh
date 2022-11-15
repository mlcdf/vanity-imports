go install github.com/cespare/reflex@latest

reflex \
    --regex='\.(go|html)$' \
    --inverse-regex='dist\/.*\.html$' \
    --start-service \
    -- sh -c 'go run . && python3 -m http.server 8000 --directory dist'