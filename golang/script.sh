#!/bin/bash

for dir in */; do
    if [[ -d "$dir" ]]; then
        new_name="${dir//-/_}"
        new_name="${new_name%/}" # Remove trailing slash for mv
        if [[ "$dir" != "$new_name/" ]]; then
            mv "$dir" "$new_name"
            echo "Renamed: $dir -> $new_name"
        fi
    fi
done

