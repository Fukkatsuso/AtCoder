n = gets.chomp.to_i
s = gets.chomp.chars

c = s[0]
ans = 1

(1..n-1).each do |i|
  if c != s[i]
    c = s[i]
    ans += 1
  end
end

puts ans