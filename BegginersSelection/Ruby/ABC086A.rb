nums = gets.chomp.split(" ").map(&:to_i)
if nums[0] % 2 == 0 || nums[1] % 2 == 0
    puts "Even"
else
    puts "Odd"
end