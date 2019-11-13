line = gets.chomp.split(" ").map(&:to_i)
n = line[0]
y = line[1]
ans = [-1, -1, -1]
(0..n).each do |a|
    (0..n-a).each do |b|
        c = n - a - b
        sum = a*10000 + b*5000 + c*1000
        if sum == y
            ans[0] = a
            ans[1] = b
            ans[2] = c
        end
    end
end

puts "#{ans[0]} #{ans[1]} #{ans[2]}"