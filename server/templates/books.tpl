<!DOCTYPE html>
<html lang="en">
<head>
   <meta charset="UTF-8">
   <meta name="viewport" content="width=device-width, initial-scale=1.0">
   <title>Gopher Library</title>
   <link href="https://unpkg.com/tailwindcss@^1.0/dist/tailwind.min.css" rel="stylesheet">
   <link rel="stylesheet" href="/static/css/books.css">
</head>
<body>
   <div class="box mx-auto p-8 border border-solid border-white">
      <h1 class="text-3xl text-center">Your Library</h1>
      <div class="book-list">
         {{ range . }}
         <div class="book-row border-b border-opacity-25 border-solid border-gray-300 mb-4">
            <p> <span class="text-sm text-opacity-50">title: </span> {{ .Title }} <span class="action text-xs float-right"><a href="">edit</a> | <a href="">delete</a></span></p>
            <p> <span class="text-sm text-opacity-50">author: </span> {{ .Author }}</p>
            <p> <span class="text-sm text-opacity-50">id: </span> {{ .ID }}</p>
         </div>
         {{ end }}
      </div>
   </div>
</body>
</html>