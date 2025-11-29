This project is a self directed exercise and part of the boot.dev Golang curriculum  

To run this project, you’ll need PostgreSQL installed and running on your machine, as well as a working installation of Go (version 1.21 or later). Make sure your PATH is configured so the go command is available in your terminal.

Installing the gator CLI

You can install the gator command-line tool using Go’s built-in install command. From the project root, run:

go install ./cmd/gator

This will place the compiled gator binary in your Go bin directory (typically $GOPATH/bin or $HOME/go/bin). Ensure that directory is on your PATH so you can run gator from anywhere.

Create a .gatorconfig.json file $HOME directory and give it fields for db_url, current_user_name, and connection_string.
My local test server config file contents look like this for example
{"db_url":"postgres://example","current_user_name":"kahya","connection_string":"postgres://postgres:password@localhost:5432/gator?sslmode=disable"}

You can run the program with commands with the go run . <command> <args>
The available commands for example are
register "username" - adds a user
users - shows the current users in the console
addfeed "feed_name" "feed_url" - adds an RSS feed for the current user
feeds - shows the current feeds for a user
follow "feed_url" - add feed results from a browsed feed to get posts when paired with the agg and command
following - "shows the current feed follows for a user
unfollow "feed_url" - delete a feed follow for the current user
agg - collects posts from added feeds. will refresh oldest RSS feed every minute.
browse "number_of_posts" - see most recent collected posts