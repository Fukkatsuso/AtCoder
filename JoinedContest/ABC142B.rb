n, k = gets.chomp.split(" ").map(&:to_i)
h = gets.chomp.split(" ").map(&:to_i)

ans = 0
(0..n-1).each do |i|
  ans += 1 if h[i] >= k
end

puts ans