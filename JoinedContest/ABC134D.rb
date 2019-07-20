# 謎のRE

n = gets.chomp.to_i
a = [0] + gets.chomp.split(" ").map(&:to_i)
b = [0] * (n+1)

m = []
(1..n).each do |i|
    i = n - i + 1
    j = i+i
    sum = 0
    while j <= n do
        sum ^= b[j]
        j += i
    end
    if (a[i] ^ sum) == 1
        m.push(i)
        b[i] = 1
    end
end

s = m.size
puts s
return if s == 0
m.each do |num|
    puts num
end