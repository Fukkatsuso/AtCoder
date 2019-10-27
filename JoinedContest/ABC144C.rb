n = gets.chomp.to_i
max = (n ** 0.5).to_i

ans = n - 1

(0..max-1).each do |i|
  i = max - i
  if n % i == 0
    x = i + n / i - 2
    ans = x if x < ans
  end
end

puts ans