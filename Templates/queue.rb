class SimpleQueue < Array
    alias_method :enq, :push
    alias_method :deq, :shift
end

q = SimpleQueue.new

data = [1, 2, 3, 4, 5]

data.each do |d|
    q.enq(d)
end

puts q.deq
puts q.deq