def is_sqnum(n)
    return ((n**0.5).to_i)**2 == n
end

n, d = gets.chomp.split(" ").map(&:to_i)
x = []
(0..n-1).each do |i|
    x.push(gets.chomp.split(" ").map(&:to_i))
end

sum = 0
(0..n-1).each do |i|
    (i..n-1).each do |j|
        next if i == j
        dist = 0
        (0..d-1).each do |k|
            dist += (x[i][k] - x[j][k]) ** 2
        end
        sum += 1 if is_sqnum(dist)
    end
end

puts sum