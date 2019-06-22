n, l = gets.chomp.split(" ").map(&:to_i)

sum = 0
tastes = []
(0..n-1).each do |i|
    tastes << (l + i)
    sum += (l + i)
end

max = tastes[n-1]
min = tastes[0]

if min > 0
    puts sum - min
elsif max < 0
    puts sum - max
else 
    puts sum
end