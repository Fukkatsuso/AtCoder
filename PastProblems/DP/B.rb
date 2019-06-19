n, k = gets.chomp.split(" ").map(&:to_i)
h = gets.chomp.split(" ").map(&:to_i)

# cost[i] : i-1段目到達の最小コスト
cost = []
cost[0] = 0
cost[1] = (h[1] - h[0]).abs

(2..(n-1)).each do |i|
    min = Float::INFINITY
    (1..k).each do |j|
        break if i < j
        c = cost[i-j] + (h[i]-h[i-j]).abs
        min = c if min > c
    end
    cost[i] = min
end

puts cost[n-1]