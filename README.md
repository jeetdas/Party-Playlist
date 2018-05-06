# Party Playlist

A web application that allows sharing of YouTube playlists, with a simple and easy to use UI. The application implements real-time sharing of playlists across devices and allows multiple users to add and update items together. The application is implemented using sockets on a Go infastructure.

## Usage

1. Site allows creation of a shared playlist and keeps track of who the admin is using a cookie.

2. Playlists can be shared with a simple code.

3. Shared playlists can be joined by any number of users, who can add new videos and vote on the next song.

### To run application
Run 'go run main.go' in src/ of the application directory

## Implementation

### Database (PostgreSQL):
Playlist metadata is stored in PostgreSQL (sharable code, admin cookie information, etc).
Individual playlist information is stored in another table (songs in the playlist, votes, etc).

### Backend (Go):
Maintains connections between sockets and passes updates along as they come in.
Stores information in the databse.

### Frontend (HTML + CSS + JS)
Interface to YouTube video and search API