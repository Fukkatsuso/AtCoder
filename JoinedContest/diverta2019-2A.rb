n, k = gets.chomp.split(" ").map(&:to_i)

min = 1
max = n - k + 1

puts max - min if k > 1
puts 0 if k == 1