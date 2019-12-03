n = gets.chomp.to_i
a = []
(1..2).each do |i|
    a.push(gets.chomp.split(" ").map(&:to_i))
end

max = 0
(0..n-1).each do |i|
    top = a[0][0..i].inject(:+)
    btm = a[1][i..n-1].inject(:+)
    max = top + btm if max < (top + btm)
end

puts max