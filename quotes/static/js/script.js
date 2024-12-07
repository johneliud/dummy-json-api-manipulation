let quotes = [];
let currentIndex = 0;

// fetchQuote uses a provided endpoint sent by the backend to fetch JSON encoded data.
async function fetchQuotes() {
  try {
    const response = await fetch('/quotes');
    quotes = await response.json();
    displayQuote();
  } catch (error) {
    console.error('Error fetching quotes:', error);
  }
}

// displayQuote updates the HTML with the current quote.
function displayQuote() {
  if (quotes.length === 0) return;

  const currentQuote = quotes[currentIndex];
  const quoteElem = document.getElementById('quote');
  const authorElem = document.getElementById('author');

  quoteElem.classList.remove('show');
  authorElem.classList.remove('show');

  setTimeout(() => {
    quoteElem.textContent = currentQuote.quote;
    authorElem.textContent = `— ${currentQuote.author}`;

    quoteElem.classList.add('show');
    authorElem.classList.add('show');

    console.log(
      `Quote: "${currentQuote.quote}"
Author: ${currentQuote.author}`
    );
  }, 500);

  currentIndex = (currentIndex + 1) % quotes.length;
}

// Displays a new quote after every 10 seconds.
function startSlideshow() {
  setInterval(displayQuote, 10000);
}

fetchQuotes().then(startSlideshow);
