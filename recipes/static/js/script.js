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
  const halfStar = rating % 1 >= 0.5;
  let stars = "";

  for (let i = 0; i < fullStars; i++) {
    stars += `<ion-icon name="star"></ion-icon>`;
  }

  if (halfStar) {
    stars += `<ion-icon name="star-half"></ion-icon>`;
  }

  const emptyStars = 5 - Math.ceil(rating);
  for (let i = 0; i < emptyStars; i++) {
    stars += '<ion-icon name="star-outline"></ion-icon>';
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
