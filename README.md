# feed2probe
HTTP Status Code testing utility for large numbers of URL/domains

### Requirements:
Go version => go1.13.8 

### Installation
`go get -u github.com/martgil/feed2probe`

### Features
- Multiple http codes filtering using `-f` flag
- Displaying response size for better filtering (use grep for filtering since there was no size filter ATM. Example: `cat domains.txt | feed2probe -p | grep <size-to-exclude-from-list>)`
- Probe domains using `-p` flag (Leave it blank for URL)

### Usage
Scan domains/endpoints 
```bash
cat target.txt | feed2probe 
```
Apply HTTP Code filters using `-f <status-code>` 
```bash
cat target.txt | feed2probe -f 200
# use ',' to apply multiple http code to filter
cat target.txt | feed2probe -f 200,403
```
Added `-p` for checking alive host
```bash
cat target.txt | feed2probe -p
# apply filters on when probing hosts
cat target.txt | feed2probe -f 200,302,403 -p
```

### Missing Features
- Output as screenshot for lazy recon (Possible selenium)
- Filter Output based on Content Length
