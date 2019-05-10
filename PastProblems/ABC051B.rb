max, sum = gets.chomp.split(" ").map(&:to_i)
kind = 0
(0..max).each do |x|
    (0..max).each do |y|
        z = sum - x - y
        kind += 1 if 0 <= z && z <= max
    end
end
puts kind