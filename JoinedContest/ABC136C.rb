n = gets.chomp.to_i
h = gets.chomp.split(" ").map(&:to_i)

can = true
if n > 1
    (1..n-1).each do |i|
        if h[i-1] > h[i]
            can = false
        elsif h[i-1] < h[i]
            h[i] -= 1
        end
    end
end

puts can ? "Yes" : "No"