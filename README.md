# Disperse

Disperse is a utility to aid distribution of content through IPFS by taking advantage of how IPFS stores content. Unless otherwise configured, whenever you request content from a public gateway, it remains in the gateway's cache until it is garbage collected; The same should in theory hold true for requesting content from other IPFS nodes.

Using Disperse, one can request a given CID from all known public gateways, and perhaps in time from other IPFS nodes. This can aid with making content more readily available throughout the network increasing availability, as well as retrieval times. This project is a work in progress.
