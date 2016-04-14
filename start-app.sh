# Start and watch

echo "Starting server..."
go run *.go &

echo "Building frontend script..."
cd _static
npm run watch &
cd ..

echo "Start server watcher..."
inotifywait -mqr -e close_write "main.go" "routes.go" "./apps/" "./core/" "./conf/" | while read line
do
    echo "Rebuilding Server... "
    go run *.go &
done
