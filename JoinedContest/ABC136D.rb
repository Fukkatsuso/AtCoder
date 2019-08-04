s = gets.chomp.split("").map{|i| i=="R" ? 1 : 0}
n = s.length
ans = [0] * n

(0..n-1).each do |i|
    step = 0
    start = i
    now = i
    while s[start] ^ s[now] == 0
        now = (s[start] == 1) ? now+1 : now-1
        step += 1
    end

    if step % 2 == 0
        ans[now] += 1
    else
        if s[start] == 1
            ans[now-1] += 1
        else
            ans[now+1] += 1
        end
    end
end

puts ans.join(" ")