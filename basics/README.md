# Envoy playground - basics

Create a basic playground simulating some service-to-service communication scenarious.

```bash
for _ in {1..1000}; do curl --silent 0.0.0.0:44823; done | sort | uniq -w 8 -c
```

