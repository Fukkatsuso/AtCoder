n = gets.chomp.to_i
a = gets.chomp.split(" ").map(&:to_i).sort

if a[0] == 0 
  puts(0)
  return
end

ans = 1

a.each do |x|
  ans *= x
  if ans > 1000000000000000000
    puts(-1)
    return
  end
end

puts(ans)

