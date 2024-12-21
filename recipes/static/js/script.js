const links = document.querySelectorAll(".link");

links.forEach((li) => {
  li.addEventListener("click", () => {
    links.forEach((link) => link.classList.remove("active"));
    li.classList.add("active");
  });
});
