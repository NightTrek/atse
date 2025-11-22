import yaml
import sys
import json

# Simple YAML reader key extractor
# Usage: python3 read_yaml.py <file> <path.to.key>

def get_value(data, path):
    parts = path.split('.')
    curr = data
    for part in parts:
        if '[' in part and ']' in part:
            key = part[:part.find('[')]
            idx_str = part[part.find('[')+1:part.find(']')]
            
            # If empty brackets [], return the whole list (if key provided) or fail
            if idx_str == "":
                if key:
                    curr = curr[key]
                # If no key and empty brackets, it means root list, just return it
                continue
            
            idx = int(idx_str)
            if key:
                curr = curr[key][idx]
            else:
                # Root is list
                curr = curr[idx]
        else:
            curr = curr[part]
    return curr

if __name__ == "__main__":
    if len(sys.argv) < 3:
        print("Usage: read_yaml.py <file> <key>")
        sys.exit(1)
        
    with open(sys.argv[1], 'r') as f:
        data = yaml.safe_load(f)
        
    path = sys.argv[2]
    
    try:
        val = get_value(data, path)
        if isinstance(val, list):
             print(" ".join(str(v) for v in val))
        elif isinstance(val, bool):
             print(str(val).lower())
        else:
             print(str(val))
    except (KeyError, IndexError, TypeError) as e:
        # sys.stderr.write(f"Error reading {path}: {e}\n")
        # Only print empty string on error to behave like yq/jq optional
        print("")
