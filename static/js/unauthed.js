document.addEventListener("DOMContentLoaded", function () {
  const signinLink = document.querySelector("#sign-in");

  signinLink.addEventListener("click", () => {
    signin();
  });
});
