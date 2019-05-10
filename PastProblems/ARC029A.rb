n = gets.chomp.to_i
t = []
(0..n-1).each do
    t.push(gets.chomp.to_i)
end
t = t.sort.reverse

one = 0
two = 0
(0..n-1).each do |i|
    if one <= two
        one += t[i]
    else
        two += t[i]
    end
end

if one >= two
    puts one
else
    puts two
end