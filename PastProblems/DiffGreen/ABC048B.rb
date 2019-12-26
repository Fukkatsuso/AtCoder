line = gets.chomp.split(" ").map(&:to_i)
min = line[0]
max = line[1]
div = line[2]

# ~min : min/div + 1
# ~max : max/div + 1
ans = max/div - min/div
ans += 1 if min % div == 0
puts ans