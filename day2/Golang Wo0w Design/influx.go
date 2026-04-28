
package main

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// InfluxDB config
const (
	token  = "your-token"
	org    = "your-org"
	bucket = "your-bucket"
	url    = "http://localhost:8086"
)

func main() {
	// Create client
	client := influxdb2.NewClient(url, token)
	defer client.Close()

	writeAPI := client.WriteAPIBlocking(org, bucket)

	// Create data point
	p := influxdb2.NewPoint(
		"temperature_sensor",
		map[string]string{"location": "plant1"},
		map[string]interface{}{"value": 25.3},
		time.Now(),
	)

	// Write to InfluxDB
	if err := writeAPI.WritePoint(context.Background(), p); err != nil {
		fmt.Println("Write error:", err)
	} else {
		fmt.Println("✅ Data written successfully")
	}

	// Query data
	query := `from(bucket:"` + bucket + `")
	|> range(start: -1h)
	|> filter(fn: (r) => r._measurement == "temperature_sensor")`

	result, err := client.QueryAPI(org).Query(context.Background(), query)
	if err != nil {
		fmt.Println("Query error:", err)
		return
	}

	// Read query result
	for result.Next() {
		fmt.Printf("⏱ %s: %v°C @ %s\n",
			result.Record().Field(),
			result.Record().Value(),
			result.Record().Time().Format(time.RFC3339))
	}
	if result.Err() != nil {
		fmt.Println("Result error:", result.Err())
	}
}
