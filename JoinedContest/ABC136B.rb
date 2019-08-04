def digit_num(n)
    i = 0
    while n > 0
        i += 1
        n /= 10
    end
    return i
end

n = gets.chomp.to_i
sum = 0

(1..n).each do |i|
    sum += 1 if digit_num(i) % 2 == 1
    i += 1
end

puts sum