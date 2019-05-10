n = gets.chomp.to_i
mochis = Array.new
(0..n-1).each do |i|
    mochis[i] = gets.chomp.to_i
end
mochis = mochis.sort.uniq.reverse

puts mochis.size