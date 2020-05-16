inotifywait -r -q -e close_write -m . |
while read -r directory events filename; do
  [ "${filename##*.}" = "go" ] && go $1 $directory
done
