n = gets.chomp.to_i
hn = gets.chomp.split(" ").map(&:to_i)

cnt = [0] * n

(1..n-1).each do |i|
	i = n - 1 - i
	cnt[i] += (cnt[i+1] + 1) if hn[i] >= hn[i+1]
end

puts cnt.max