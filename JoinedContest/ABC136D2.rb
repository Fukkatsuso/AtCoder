s = gets.chomp.split("").map{|i| i=="R" ? 1 : 0}
n = s.length
ans = [0] * n
sum_sub = [0] * 2

start = 0
now = 0
step = 0
(0..n-1).each do |i|
    sum_sub[step % 2] += 1
    now += 1
    step += 1

    if now == n || (s[start] ^ s[now]) == 1
        if s[start] == 1 #右に向かう集団
            ans[now] += sum_sub[1]
            ans[now-1] += sum_sub[0]
        else
            ans[start-1] += sum_sub[1]
            ans[start] += sum_sub[0]
        end
        
        start = i+1
        now = i+1
        step = 0
        sum_sub = [0] * 2
    end
    break if start >= n
end

puts ans.join(" ")