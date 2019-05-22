# 解けず

n = gets.chomp.to_i
arr = []
(n-1).times do
    arr.push(gets.chomp.split(" ").map(&:to_i))
end
ans = Array.new(n, 0)

(0..n-2).each do |i|
    if arr[i][2] % 2 == 0
        ans[arr[i][0]-1] = 1
        ans[arr[i][1]-1] = 1
    end
end
(0..n-1).each do |i|
    puts ans[i]
end