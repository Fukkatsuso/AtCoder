n = gets.chomp.to_i
h = []
n.times do
    h << gets.chomp.split(" ").map(&:to_i)
end

# 前日までの累積の最大値, 
dp = []  # [累積, 前日の活動]

dp[0] = [h[0][0], h[0][1], h[0][2]]

(1..n-1).each do |i|
    pre = dp[i-1]
    dp[i] = [h[i][0], h[i][1], h[i][2]]
    dp[i][0] += (pre[1] > pre[2] ?  pre[1] : pre[2])
    dp[i][1] += (pre[0] > pre[2] ?  pre[0] : pre[2])
    dp[i][2] += (pre[1] > pre[0] ?  pre[1] : pre[0])
end

puts dp[n-1].max