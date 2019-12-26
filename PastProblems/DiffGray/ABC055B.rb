n = gets.chomp.to_i
pow = 1
(1..n).each do |i|
    pow = (pow * i) % (1000000007)
end
puts pow