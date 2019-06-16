n, k = gets.chomp.split(" ").map(&:to_i)
a = gets.chomp.split(" ").map(&:to_i)
cnt = 0

# しゃくとり法 AC
right = 0
sum = 0
(0..(n-1)).each do |left|
    while right < n && sum + a[right] < k do
        sum += a[right]
        right += 1
    end
    cnt += (n - right)

    if right == left
        right += 1
    else 
        sum -= a[left]
    end

end

puts cnt