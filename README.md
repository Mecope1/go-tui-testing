# go-tui-testing
Using this repo as a sandbox to explore how best to use the Termui library for a CLI chat application


## Notes on Termui for a Chat Application

~~Scrolling using the arrow keys and Page Up/Down keys seems sluggish/unresponsive. When the events do trigger, the effect is immediate though.~~

I had thought it was unresponsive, but actually I had misunderstood what the List.ScrollUp() and List.ScrollDown() 
functions were doing. Because the List structure is a slice of strings, the result of the ScrollDown() function 
is just to move a cursor forward to the next slice item, like from an index of 4 to 5. Also, this cursor position 
doesn't automatically move when an item is added to the slice. The result is that if your rendered frame was showing 
items 0-20, and you wanted to see item 21, you would have to click the down arrow key 20 times to move the cursor 
before it reached the end of the frame! You would then need to advance it one more time to shift the frame down to 
display item 21. However, the List.ScrollPageDown() function will move the cursor by a frame's length + 1. So for 
our example, it would move the cursor immediately to item 21 from 0. The next time it was triggered, it would shift to
item 42. Based on this, it may be helpful to users to change what happens when they use the arrow keys. To give the 
most intuitive experience for using arrow keys, you could first trigger the ScrollPageDown() function, and then have 
subsequent key presses in the same direction use ScrollDown(). However, this could make it awkward if you wanted to 
land exactly on an index. Any time you changed direction, it would force the cursor to go a much further distance 
than would be helpful. This could potentially make it impossible to reach certain indices if the list was smaller 
than ~2 frame lengths as well.