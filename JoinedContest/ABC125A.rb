input = gets.chomp.split(" ")
period = input[0].to_i
bsc = input[1].to_i
t = input[2].to_i

tm = 0
sum = 0
while tm <= (t + 0.5) do
    sum += bsc if (tm % period == 0) && (tm != 0)
    tm += 1
end

puts sum