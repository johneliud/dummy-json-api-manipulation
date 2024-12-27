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
