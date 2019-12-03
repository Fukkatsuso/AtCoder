n = gets.chomp.to_i
turn_left = Array.new(n, 0)
turn_right = Array.new(n, 0)
s = gets.chomp.split("")

(1..n-1).each do |i| 
    if s[i-1] == "W"
        turn_left[i] = (turn_left[i-1]+1)
    else
        turn_left[i] = (turn_left[i-1])
    end
end
(0..n-2).each do |i|
    i = n-2 - i
    if s[i+1] == "E"
        turn_right[i] = (turn_right[i+1]+1)
    else 
        turn_right[i] = turn_right[i+1]
    end
end

min = n
(0..n-1).each do |i|
    sum = turn_left[i] + turn_right[i]
    if min > sum
        min = sum
    end
end
puts min