# ASCII Art Web Project

## Description
This project generates ASCII art using user-provided text and a selected banner style. The web interface allows users to input text and choose from available banners to generate ASCII representations.

## Author
- Youssef JAOUHAR

## Usage
1. Clone the repository.
2. Place the required banner files (`standard.txt`, `shadow.txt`, `thinkertoy.txt`) in the `Banners/` directory.
3. To run the server:
   go run main.go

## Algorithm Overview
1. Server Initialization:

The server is set up to listen on port 8080.
Two routes are registered:
/: Displays the homepage where users can submit their text and banner selection.
/ascii-art: Handles the form submission to generate ASCII art.
2. Home Page (/):

The HomeHandler serves the homepage by rendering an HTML template (index.html).
It checks if the URL path is exactly /. If not, a 404 error is returned.
The template is parsed, and empty data (no text or ASCII art) is passed for initial rendering.
3. ASCII Art Generation (/ascii-art):

The AsciiArtHandler processes form submissions via POST requests.
It parses the form data, validates the input text (ensuring it’s non-empty and under 1000 characters), and retrieves the selected banner (defaulting to "standard").
The corresponding banner data is read from a file.
ASCII art is generated from the input text using the selected banner.
The result (ASCII art) is rendered into the same HTML template and returned to the user.
4. Error Handling:

Various HTTP status codes are returned for error scenarios:
200 OK for successful requests.
404 Not Found if the URL path is invalid.
400 Bad Request for invalid input data.
500 Internal Server Error for template or banner reading errors.
5. Template Rendering:

HTML templates are used to structure the web pages. The form data and generated ASCII art are passed into the templates for rendering the output to the user.