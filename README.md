# feed2probe
HTTP Status Code testing utilility for large numbers of URL/domains

### Requirements:
Go version => go1.13.8 

### Usage
Scan domains/endpoints 
```bash
cat target.txt | feed2probe 
```
Apply HTTP Code filters using `-f <status-code>` 
```bash
cat target.txt | feed2probe -f 200
```
