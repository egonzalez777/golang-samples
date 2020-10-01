## Simple algorithm to retrieve the last 5 events of a given timeframe.

### Specifications:
We have a sample slice that contains potentially duplicate entries. The task at hand is to retrieve the last 5 entries without having duplicates in the final list or the last viewed events

####
Sample list:

```
sampleList := []int{1, 2, 30, 40, 20, 2, 40, 5, 40, 5, 10, 5}
```

Expected result list (Last 5 values):
```
resultList = [5, 10, 40, 2, 20]
```