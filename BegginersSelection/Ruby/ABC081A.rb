s = gets.chomp.split("").map(&:to_i)
n = 0
s.each do |i|
    n += i
end
puts n