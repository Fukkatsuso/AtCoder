n = gets.chomp.to_i
can = false

(1..9).each do |i|
  if n % i == 0
    if n / i <= 9
      can = true
      break
    end
  end
end

if can
  puts "Yes"
else
  puts "No"
end