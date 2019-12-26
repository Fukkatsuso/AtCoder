n = gets.chomp.to_i
point = []
(0..n-1).each do
    point.push(gets.chomp.split(" ").map(&:to_f))
end

max = 0
(0..n-1).each do |i|
    (0..n-1).each do |j|
        dist_x = point[i][0] - point[j][0]
        dist_y = point[i][1] - point[j][1]
        dist = Math.sqrt(dist_x ** 2 + dist_y ** 2)
        max = dist if max < dist
    end
end
puts max