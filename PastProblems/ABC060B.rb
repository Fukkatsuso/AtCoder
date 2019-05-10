line = gets.chomp.split(" ").map(&:to_i)
a = line[0]
b = line[1]
c = line[2]
can = false
(1..b).each do |i|
    can = true if a * i % b == c
end

if can
    puts "YES"
else 
    puts "NO"
end