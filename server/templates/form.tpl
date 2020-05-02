<!DOCTYPE html>
<html lang="en">
<head>
   <meta charset="UTF-8">
   <meta name="viewport" content="width=device-width, initial-scale=1.0">
   <title>Gopher Library - New Book</title>
   <link href="https://unpkg.com/tailwindcss@^1.0/dist/tailwind.min.css" rel="stylesheet">
   <link rel="stylesheet" href="/static/css/form.css">
</head>
<body>
   <div class="box mx-auto p-8 border border-solid border-white">
      {{if .Book }}
      <form class="text-center" action="{{ .Action }}" method="POST">
         <label class="block text-sm" for="title">Title</label>
         <input type="hidden" name="_method" value="{{ .Method }}">
         <input class="mb-4 mt-2" type="text" value="{{ .Book.Title }}" name="title" id="title">
         <label class="block text-sm" for="author">Author</label>
         <input class="mb-4 mt-2" type="text" value="{{ .Book.Author }}" name="author" id="author">
         <button class="block ml-32 mb-4 text-gray-100 w-24 rounded-md h-8 border-2 border-yellow-400 border-solid" type="submit">submit</button>
      </form>
      {{ else }}
      <form class="text-center" action="/books/" method="POST">
         <label class="block text-sm" for="title">Title</label>
         <input class="mb-4 mt-2" type="text" name="title" id="title">
         <label class="block text-sm" for="author">Author</label>
         <input class="mb-4 mt-2" type="text" name="author" id="author">
         <button class="block ml-32 mb-4 text-gray-100 w-24 rounded-md h-8 border-2 border-yellow-400 border-solid" type="submit">submit</button>
      </form>
      {{ end }}
   </div>
</body>
</html>