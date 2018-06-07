The idea behind the Observer pattern is simple - to subscribe to some event that will trigger some behavior on many 
subscribed types. So we uncouple an event from its possible handlers.

For example, imagine a login button. We could code that when the user clicks the button, the button color changes, an 
action is executed, and a form check is performed in the background. But with the Observer pattern, the type that 
changes the color will subscribe to the event of the clicking of the button. The type that checks the form and the type
that performs an action will subscribe to this event too.