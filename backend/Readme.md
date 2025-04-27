# Apica Search - Backend Service

This is a backend service built in Go. It allows you to upload Parquet files, store records, and search through them efficiently. It uses Gin for the web server and a simple in-memory index to quickly find matching records.

## How It Works

### Key Parts:

1. **Handlers**:
   - **SearchHandler**: Handles search requests. It checks the query and returns matching records.
   - **UploadAndStore**: Handles file uploads. It saves Parquet files, parses them into records, and stores them for searching.

2. **Parser**:
   - **LoadParquetFile**: Reads and parses Parquet files into structured data that can be used in searches.

3. **Search**:
   - Takes a search query, breaks it into words, and looks up each word in an index to find matching records.

4. **Record Store (Utils)**:
   - **RecordStore**: Stores all the records and indexes them for fast search.
   - It uses an inverted index, meaning it keeps track of which records contain which words, so searching is quick.

5. **Main**:
   - The main file runs the app. It loads all the Parquet files when the app starts and sets up the web server to handle requests.

### Design Choices:
- **Concurrency**: Uses Go's `sync.RWMutex` to let multiple parts of the app read the data at the same time without messing up.
- **In-memory Index**: The records are stored in memory with an index for fast searching. This means it's really fast but needs enough memory to store all the data.
- **Scalability**: The current design works well for small to medium-sized datasets. For larger datasets, you could use something like Redis or a database.

## How to Build and Run

### Prerequisites:
1. Go 1.18+.
2. Install dependencies using Go modules.

### Steps:

1. Clone the repository:
   ```bash
   git clone <repository_url>
   cd apica-search
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the app:
   ```bash
   go build -o apica-search
   ```

4. Run the app:
   ```bash
   ./apica-search
   ```

   The server will run on `http://localhost:8181`.

### API Endpoints:

1. **POST /upload**:
   - Uploads a Parquet file and indexes it.
   - Example:
     ```bash
     curl -X POST -F "file=@path/to/file.parquet" http://localhost:8181/upload
     ```

2. **GET /search?query=<search_term>**:
   - Searches for records that match the query.
   - Example:
     ```bash
     curl "http://localhost:8181/search?query=test"
     ```

## Benchmarks and Performance

### Benchmark Results:
```bash
BenchmarkRecordStore_Set-8      	  166272	      7591 ns/op	    2673 B/op	      29 allocs/op
BenchmarkRecordStore_Search-8   	47875905	        25.23 ns/op	       0 B/op	       0 allocs/op
```

- **Record Store Set**: 166,272 operations in 10 seconds, average 7.591 microseconds per operation.
- **Record Store Search**: 47,875,905 searches in 10 seconds, average 25.23 nanoseconds per operation.

### Performance Tips:
- **Memory Usage**: The in-memory index makes searching fast, but it uses memory. If the dataset grows too big, you could look into storing data on disk or in a distributed system.
- **Parallelism**: The app is designed to use Go's concurrency well, but there could be improvements with batch processing for indexing and memory usage.

## Future Enhancements

1. **Distributed Storage**:
   - Use a distributed database or something like Redis to handle larger datasets more efficiently.

2. **Advanced Search Features**:
   - Add features like fuzzy search, relevance ranking, or support for complex queries (AND, OR conditions).

3. **Persistent Data**:
   - Store the indexed records in a database, so you don’t lose them when the app restarts.

4. **Better Benchmarking**:
   - Run more benchmarks to test bulk indexing, search performance, and memory usage.
