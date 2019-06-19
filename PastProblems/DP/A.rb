def min2(a, b)
    return a <= b ? a : b
end

n = gets.chomp.to_i
h = gets.chomp.split(" ").map(&:to_i)

# cost[i] : i-1段目到達の最小コスト
cost = Array.new(h.length+1, 0)

(1..(n-1)).each do |i|
    if i == 1
        cost[i] = (h[i]-h[i-1]).abs
        next
    end
    cost[i] = min2(cost[i-1] + (h[i]-h[i-1]).abs,
                    cost[i-2] + (h[i]-h[i-2]).abs)
end

puts cost[n-1]