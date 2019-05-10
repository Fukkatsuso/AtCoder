a, b, c, d = gets.chomp.split("").map(&:to_i)
op = []
flag = false
(0..1).each do |i|
    (0..1).each do |j|
        (0..1).each do |k|
            if a + (i*2-1)*b + (j*2-1)*c + (k*2-1)*d == 7
                op.push(i, j, k)
                flag = true
            end
            break if flag
        end
        break if flag
    end
    break if flag
end
(0..2).each do |i|
    if op[i] == 1
        op[i] = '+'
    else 
        op[i] = '-'
    end
end

puts "#{a}#{op[0]}#{b}#{op[1]}#{c}#{op[2]}#{d}=7"