n, k = gets.chomp.split(" ").map(&:to_i)
an = gets.chomp.split(" ").map(&:to_i)

cnt = 0
sub_sum = 0

# TLE
(0..(n-1)).each do |start|
    sub_sum = an[start]
    cnt += 1 if sub_sum >= k #1項のみでk以上を満たす場合
    last = start + 1
    while last > start && last < n do
        sub_sum += an[last]
        if sub_sum >= k
            cnt += (n - last)
            break
        end
        last += 1
    end
end

puts cnt