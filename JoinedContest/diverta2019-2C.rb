n = gets.chomp.to_i
an = gets.chomp.split(" ").map(&:to_i)

line = []
an = an.sort.reverse  # 降順

# until an.length == 2
(n-2).times do
    min1 = an.pop
    min2 = an.pop
    line.push([min1, min2])
    min = min1 - min2
    an.push(min)
end
line.push([an[0], an[1]])

puts an[0] - an[1]

until line.empty?
    l = line.shift
    puts "#{l[0]} #{l[1]}"
end