package priorityqueue


type node struct{
	priority int
	data int
	next *node
	prev *node
}

type priorityQueue struct{
	head *node
	tails  []*node
	nprios int
}

func(pQueue *priorityQueue) pQueue_init (nprios int){
	if(nprios <= 0){
        return
    }

    pQueue.head = nil
	pQueue.tails = make([]*node, nprios)
	pQueue.nprios = nprios
}



// Takes in a node that has a value > 0. Priority based on smallest - largest ints
func(pQueue *priorityQueue) insert (priority, data int){
	// Valid priority queue data
	if data < 0 || priority > pQueue.nprios - 1{
		return
	}

	// Create a new node
	newNode := &node{priority: priority, data: data, next: nil, prev: nil}
	
	// If list is not initiaziled


    // Check that if list is empty so node is head
    if(pQueue.head == nil){
        pQueue.head = newNode;
        pQueue.tails[priority] = newNode;
        return;
        //Check if new node will be head
    }else if(newNode.priority < pQueue.head.priority){
        pQueue.head.prev = newNode;
        newNode.next = pQueue.head;
        pQueue.head = newNode;
        pQueue.tails[priority] = newNode;
        return;
    }
	return
}

func main() {


}