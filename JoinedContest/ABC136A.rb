a, b, c = gets.chomp.split(" ").map(&:to_i)

w = c - (a - b)
w = 0 if w < 0

puts w