#!/bin/bash

# Record the start time of the entire script
start_time=$(date +%s)

# Loop through all subdirectories starting with "problem"
for dir in problem*/; do
    # Check if the solution.go file exists in the directory
    if [ -f "$dir/solution.go" ]; then
        # Record the start time for the current Go file
        file_start_time=$(date +%s)
        
        # echo "Running $dir/solution.go..."
        cd $dir
        # Run the Go solution file
        go run solution.go
        
        # Record the end time for the current Go file
        file_end_time=$(date +%s)
        
        # Calculate elapsed time for the current Go file
        elapsed_time=$((file_end_time - file_start_time))
        
        # Report the time taken to execute the current file
        echo "Time taken for $dir: $elapsed_time s"
	cd ..
    else
        echo "No solution.go found in $dir"
    fi
done

# Record the end time of the entire script
end_time=$(date +%s)

# Calculate total elapsed time for the entire script
total_elapsed_time=$((end_time - start_time))

# Report the total time taken for the entire script
echo "Total time taken for the script: $total_elapsed_time s"
