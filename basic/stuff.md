```bash
for _ in {1..1000}; do curl --silent 0.0.0.0:44823; done | sort | uniq -w 8 -c
```

