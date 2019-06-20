n, weight = gets.chomp.split(" ").map(&:to_i)
wv = []
n.times do
    wv << gets.chomp.split(" ").map(&:to_i)
end

dp = Array.new(weight+1, -1) # dp[重さ] = 価値
dp[0] = 0

wv.each do |w, v|
    (weight - w).downto(0) do |i|
        next if dp[i] < 0
        nv = dp[i] + v #価値
        nw = i + w #重さ
        dp[nw] = nv if dp[nw] < nv
    end
end

# 1つの品物は1度しか選べない
# (1..weight).each do |w|
#     max = 0
#     (0..(n-1)).each do |i|
#         next if w < wv[i][0]
#         value = dp[w-wv[i][0]] + wv[i][1]
#         max = value if max < value
#     end
#     dp[w] = max
#     end
# end

puts dp.max