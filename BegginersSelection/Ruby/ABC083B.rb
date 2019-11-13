def sum_digit(i, a, b)
    # 各桁の総和
    sum = 0
    while(i > 0) do
        sum += i % 10
        i /= 10
    end
    if a <= sum && sum <= b
        return true
    else 
        return false
    end
end

line = gets.chomp.split(" ").map(&:to_i)
n = line[0]
a = line[1]
b = line[2]
ans = 0

(1..n).each do |i|
    if sum_digit(i, a, b)
        ans += i
    end
end
puts ans