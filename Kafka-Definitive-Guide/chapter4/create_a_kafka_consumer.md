## Thread Safety

You can'v have multiple consumers that belong to the same group in one thread, and you can't have multiple threads safety use the same consumer. One consumer per thread is the rule. To run multiple consumers in the same group in one application, you will need to run each in its own thread.