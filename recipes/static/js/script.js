const links = document.querySelectorAll(".link");

links.forEach((li) => {
  li.addEventListener("click", () => {
    links.forEach((link) => link.classList.remove("active"));
    li.classList.add("active");
  });
});

let recipes = [];

const renderRatingStars = (rating) => {
  const fullStars = Math.floor(rating);
  const halfStar = rating - fullStars >= 0.5;
  let stars = "";

  for (let i = 0; i < fullStars; i++) {
    stars += `<ion-icon name="star"></ion-icon>`;
  }

  if (halfStar) {
    stars += `<ion-icon name="star-half"></ion-icon>`;
  }

  const emptyStars = 5 - Math.ceil(rating);
  for (let i = 0; i < emptyStars; i++) {
    stars += `<ion-icon name="star-outline"></ion-icon>`;
  }
  return stars;
};

const updateFeaturedRecipe = (recipe) => {
  const content = document.getElementById("featured-recipe-content");
  if (!content) {
    return;
  }

  content.innerHTML = `
  <div class="featured-image">
    <img src="${recipe.Image}" alt="${recipe.Name}" />
  </div>
        
  <div class="featured-details">
    <h2>${recipe.Name}</h2>
    <div class="featured-meta">
      <div class="ratings-container">
        ${renderRatingStars(recipe.Rating)}
      </div>
      <p class="cuisine">${recipe.Cuisine}</p>
    </div>
    <p class="description">Prep Time: ${
      recipe.PrepTimeMinutes
    } mins • Cook Time: ${recipe.CookTimeMinutes} mins</p>
  </div>
  `;
  content.onclick = () => (window.location.href = `/recipe/${recipe.ID}`);
};

const getRandomRecipe = () => {
  const randomIndex = Math.floor(Math.random() * recipes.length);
  return recipes[randomIndex];
};

document.addEventListener("DOMContentLoaded", () => {
  const recipeElements = document.querySelectorAll(".recipe");
  recipeElements.forEach((elem) => {
    const recipe = {
      ID: elem.onclick.toString().match(/\/recipe\/(\d+)/)[1],
      Name: elem.querySelector(".recipe-name").textContent,
      Image: elem.querySelector("img").src,
      Rating: parseFloat(
        elem.querySelector(".ratings-container").getAttribute("data-rating")
      ),
      Cuisine: elem.querySelector(".cuisine-container p").textContent,
      PrepTimeMinutes: parseInt(
        elem.querySelector(".prep-time").getAttribute("data-preptime")
      ),
      CookTimeMinutes: parseInt(
        elem.querySelector(".cook-time").getAttribute("data-cooktime")
      ),
    };
    recipes.push(recipe);
  });

  setInterval(() => {
    const randomRecipe = getRandomRecipe();
    updateFeaturedRecipe(randomRecipe);
  }, 15000);
});
