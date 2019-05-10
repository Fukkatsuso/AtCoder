line = gets.chomp.split(" ").map(&:to_i)
if line[0] <= line[2]
    begin_a = line[0]
    end_a   = line[1]
    begin_b = line[2]
    end_b   = line[3]
else
    begin_a = line[2]
    end_a   = line[3]
    begin_b = line[0]
    end_b   = line[1]
end

if end_a < begin_b
    time = 0
elsif end_a < end_b
    time = end_a - begin_b
else
    time = end_b - begin_b
end

puts time