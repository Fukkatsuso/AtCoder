n, k = gets.chomp.split(" ").map(&:to_i)
r, s, p = gets.chomp.split(" ").map(&:to_i)
t = gets.chomp

ans = 0
win = Array.new(n, true)
(0..n-1).each do |i|
	point = 0
	if t[i] == 'r'
		point = p
	elsif t[i] == 's'
		point = r
	elsif t[i] == 'p'
		point = s
	end

	if (i >= k) && (t[i-k] == t[i]) && win[i-k]
		point = 0
		win[i] = false
	end
	ans += point
end

puts ans
