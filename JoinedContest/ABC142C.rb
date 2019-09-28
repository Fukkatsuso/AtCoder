n = gets.chomp.to_i
a = gets.chomp.split(" ").map(&:to_i)
ans = [0] * n

(0..n-1).each do |i|
  ans[a[i] - 1] = i + 1
end

(0..n-1).each do |i|
  print ans[i]
  if i < n - 1
    print " "
  else
    puts ""
  end
end