n = gets.chomp.to_i
a = gets.chomp.split(" ").map(&:to_i).sort.reverse
alice = 0
bob = 0

(0..n-1).each do |i|
    alice += a[i] if i % 2 == 0
    bob += a[i]   if i % 2 == 1
end

puts alice - bob