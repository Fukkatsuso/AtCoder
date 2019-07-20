n, d = gets.chomp.split(" ").map(&:to_i)

rem = n
ans = 0
while rem > 0 do
    ans += 1
    rem -= (2*d + 1)
end
puts ans