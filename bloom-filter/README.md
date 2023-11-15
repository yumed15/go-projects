# Bloom Filter Implementation
A Bloom filter is a probabilistic data structure, named after Burton Howard Bloom, who invented them in the 1970s. It is based on hashing designed to tell you, rapidly and memory-efficiently, whether an element is present in a set. Elements are not added to the set but their hash is.

The price paid for this efficiency is that a Bloom filter is a probabilistic data structure: it tells us whether the element is either definitely not in the set or that it may be in the set, thus false positives are possible.

More at [Optimising Performance: The Cool Magic of Bloom Filters](https://thecraftydev.substack.com/p/optimising-performance-the-cool-magic)