stack = []
n = 5

(1..n).each do |i|
    stack.push(i)
end

p stack

n.times do
    print "#{stack.pop} "
end
puts ""