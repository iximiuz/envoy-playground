# Envoy playground - request hedging

Create a playground simulating request hedging.

## TODO:

- create Service B simulating long tail latencies
- shed some traffic to Service B and plot latency histogram
- create Service A accessing Service B 
- shed some traffic to Service B through Service A and plot
  latency histogram
- add envoy in between Service A and Service B with request
  hedging activated
- shed some traffic and show the improved latency, also
  show that the amount of extra requests to Service B is small
- degrade Service B and repeat exercise, show that we doubled
  traffic to Service B while the latency of Service A hasn't
  been improved much
