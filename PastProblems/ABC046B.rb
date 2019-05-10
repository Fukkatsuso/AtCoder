line = gets.chomp.split(" ").map(&:to_i)
n = line[0]
k = line[1]
ans = k
if n > 1
    (2..n).each do |i|
        ans *= (k - 1)
    end
end

puts ans