n = gets.chomp.to_i
vs = gets.chomp.split.map(&:to_i)
cs = gets.chomp.split.map(&:to_i)

sum = 0
(0..n-1).each do |i|
    sum += (vs[i] - cs[i]) if (vs[i] - cs[i]) > 0
end

puts sum