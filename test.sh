echo
for i in {1..500}
  do
    echo "$i"
    curl $1
    sleep 0.05
done
