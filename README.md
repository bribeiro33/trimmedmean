# trimmedmean

A Go package for computing **trimmed means** — averages that remove a proportion of the lowest and highest observations before computing the mean.

## Features

- Symmetric and asymmetric trimming support  
- Minimal dependencies  
- Works on any `[]float64` slice  
- Optional helper to convert `[]int` → `[]float64`  

## Installation

```bash
go mod init <project-name>
go get github.com/bribeiro33/trimmedmean
```

## Example

```go
package main
import (
	"fmt"
	"log"

	"github.com/bribeiro33/trimmedmean"
)

func main() {
	data := []float64{1, 2, 3, 4, 5, 100}

	// Symmetric 10% trim: remove 10% from both ends
	tmean, err := trimmedmean.TrimmedMean(data, 0.10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Symmetric 10%% trimmed mean: %.2f\n", tmean)

	// Asymmetric trim: remove 5% from low end, 20% from high end
	tmean2, err := trimmedmean.TrimmedMean(data, 0.05, 0.20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Asymmetric trimmed mean: %.2f\n", tmean2)
}
```

## Reference 

```
func TrimmedMean(data []float64, trimArgs ...float64) (float64, error)
```

Calculates the trimmed mean from a slice of float64 values.

| Parameter   | Type               | Description                                                                                                    |
| ----------- | ------------------ | -------------------------------------------------------------------------------------------------------------- |
| `data`      | `[]float64`        | The input data slice                                                                                           |
| `trimArgs`  | `float64` | Trimming proportions:<br>• One argument -> symmetric trimming<br>• Two arguments -> low/high asymmetric trimming |
| **Returns** | `(float64, error)` | The trimmed mean and an error if trimming removes all data or arguments are invalid  |

```
TrimmedMean(data, 0.10)       // symmetric 10% trim
TrimmedMean(data, 0.05, 0.2)  // 5% low trim, 20% high trim
```

```func ToFloat64(ints []int) []float64```

Converts a slice of ints to a slice of float64 values.
```
ints := []int{1, 2, 3, 4, 5}
floats := trimmedmean.ToFloat64(ints)
// []float64{1, 2, 3, 4, 5}
```

## More
See trimmedmean-demo github.com/bribeiro33/trimmedmean-demo for full demo and validation


