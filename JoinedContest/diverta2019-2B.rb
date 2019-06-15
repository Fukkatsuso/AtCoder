n = gets.chomp.to_i
balls = []
dist_table = []

n.times do
    balls.push(gets.chomp.split(" ").map(&:to_i))
end

# a->bへの距離をメモ
balls.each_index do |a|  # 最初に選ぶボール
    dists = []
    balls.each_index do |b|
        next if a >= b
        dists.push([balls[b][0]-balls[a][0], balls[b][1]-balls[a][1]])
    end
    dist_table.push(dists) if dists != []
end

cost_min = n

zero = 0
dist_table.each do |t|
    dist_table.each do |u|
        print "#{t} : #{u}"
        zero += 1 if t == u
        puts "->#{zero}"
    end
    cost = n - zero
    cost_min = cost if cost_min > cost
end

p cost_min
