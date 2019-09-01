a, b = gets.chomp.split(" ").map(&:to_i)

if b == 1
	puts 0
else
	tap = 1
	while true do
		break if tap * a - (tap - 1) >= b
		tap += 1
	end
	puts tap
end
