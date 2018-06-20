Sometimes it's useful to begin accepting work for your pipeline even though the pipeline is not yet ready for more.
This process is called *queuing*.

All this means is that once your stage has completed some work, it stores it in a temporary location in memory so that
other stages can retrieve it later, and your stage doesn't need to hold a reference to it.