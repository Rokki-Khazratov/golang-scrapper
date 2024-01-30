# GoScraper: E-commerce Data Scraper and API BUILDING with Gorilla Mux

GoScraper is a powerful data scraper written in Golang, utilizing the Gorilla Mux library for seamless RESTful API development. This application is designed to scrape data from popular e-commerce websites and Apple dealers in Central Asia. It goes beyond the typical scraping process by extracting information directly from HTML sources, allowing you to build APIs and retrieve images.

## Features

- **Golang Powered:** Developed entirely in Golang for efficiency, speed, and ease of use.
- **Gorilla Mux:** Utilizes the Gorilla Mux router for creating robust and scalable RESTful APIs.
- **E-commerce Data Scraper:** Scrapes data from popular e-commerce platforms and Apple dealers in Central Asia.
- **HTML Parsing:** Extracts data directly from HTML, providing accurate and up-to-date information.
- **Image Retrieval:** Fetches images associated with the scraped data, enhancing the overall value of the API.

## Thumbnail : 

![thumbnail](https://github.com/Rokki-Khazratov/web-scrapper-to-api/assets/93393135/8dbfd907-08df-4f89-9a71-8c8991ee6bb4)


## Getting Started

Follow these instructions to get the project up and running on your local machine.

### Prerequisites

- [Golang](https://golang.org/dl/) installed on your system.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Rokki-Khazratov/web-scrapper-to-api.git

2. Navigate to the project directory:

   ```bash
   cd goscraper

3. Run the main.go file:
   ```bash
   go run main.go

# Parse data from collections/
/api/parse: Sends a POST request with the following JSON payload to initiate scraping for a specific URL:

```json
{
  "url": "http://brostore.uz/collections/...."
}
```

Example Response:
```json
{
  "id": 1,
  "title": "Some",
  "variant-title": "some",
  "image-1": "link",
  "image-2": "link"
}
```

## Videos : 


frontend example
![front](https://github.com/Rokki-Khazratov/web-scrapper-to-api/assets/93393135/f233518f-3d44-4929-b148-60f38342c9ac)


postman api
![postman](https://github.com/Rokki-Khazratov/web-scrapper-to-api/assets/93393135/3f2d6f48-4506-4209-818b-3ef9400e615d)



This command will start the application, and you should see output indicating that the server is running.

Open your web browser and access the API at http://localhost:8080.
