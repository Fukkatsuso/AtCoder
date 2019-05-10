n = gets.chomp.to_i
t = [0, 0]
x = [0, 0]
y = [0, 0]

can = true
(1..n).each do |i|
    line = gets.chomp.split(" ").map(&:to_i)
    t = [t[1], line[0]]
    x = [x[1], line[1]]
    y = [y[1], line[2]]
    dt = t[1]-t[0]
    dist = (x[1]-x[0]).abs + (y[1]-y[0]).abs
    if dt < dist
        can = false
    elsif dt % 2 != dist % 2
        can = false
    end
end

if can
    puts "Yes"
else
    puts "No"
end