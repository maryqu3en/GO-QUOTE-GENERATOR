const quote = document.getElementById("quote");
const author = document.getElementById("author");
const bookTitle = document.getElementById("book-title");
const button = document.getElementById("button");
const URL = "http://localhost:3030";

button.addEventListener("click", async () => {
  try {
    const response = await fetch(`${URL}/api/quotes/random`);
    const data = await response.json();

    quote.innerHTML = data.quote;
    author.innerHTML = data.author;
    bookTitle.innerHTML = data.book;
  } catch (err) {
    alert(err.message);
  }
});
