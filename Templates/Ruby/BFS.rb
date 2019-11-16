class SimpleQueue < Array
    alias_method :enq, :push
    alias_method :deq, :shift
end


def bfs(start, target)
    queue = SimpleQueue.new
    queue.enq(start)

    until queue.empty?
        new_node = queue.deq
        if new_node.data == target.data
        end
        unless new_node.is_visited
            new_node.visit
            new_node.nodes.each do |n|
                queue.enq(n) unless n.is_visited
            end
        end
    end
end