# TLE!

def product(first, diff, n)
    ret = first
    (1..n-1).each do |i|
        ret = (ret * ((first + diff * i) % 1000003)) % 1000003
    end
    return ret
end


q = gets.chomp.to_i
ans = Array.new

q.times do
    input = gets.chomp.split(" ").map(&:to_i)
    ans.push(product(input[0], input[1], input[2]))
end

(0..q-1).each do |i|
    puts ans[i]
end