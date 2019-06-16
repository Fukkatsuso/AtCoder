n, x = gets.chomp.split(" ").map(&:to_i)
l = gets.chomp.split(" ").map(&:to_i)

d = 0
cnt = 1

(0..n-1).each do |i|
    d = d + l[i]
    cnt += 1 if d <= x
end

puts cnt