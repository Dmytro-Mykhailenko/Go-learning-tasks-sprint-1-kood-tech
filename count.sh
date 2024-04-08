count=$(find . -type f -o -type d | wc -l)

result=$((count * 5))

printf "\n\tTotal files * 5: %d\v\n" "$result"

