iterations=$1

if [ $iterations -gt 100 ]; then
    iterations=100
fi

for ((i = 1; i <= $iterations; i++)); do
    echo "$i times I've printed dmytromykhailenko"
done

