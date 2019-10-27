n = gets.chomp.to_i
d = gets.chomp.split(" ").map(&:to_i)

ans = 0

(0..n-1).each do |i|
  (0..n-1).each do |j|
    next if i == j
    ans += d[i] * d[j]
  end
end

puts ans / 2