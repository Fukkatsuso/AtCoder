n = gets.chomp.to_i
jobs = []
n.times do
    jobs << gets.chomp.split(" ").map(&:to_i)
end

# 締め切りが早い順にソート
jobs = jobs.sort_by{|x| x[1]}

time = 0
can = true

jobs.each do |t, lim|
    time += t
    if time > lim
        can = false
        break
    end
end

puts 'Yes' if can
puts 'No' if !can