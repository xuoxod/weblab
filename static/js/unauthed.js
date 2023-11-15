document.addEventListener("DOMContentLoaded", function () {
  const signinLink = document.querySelector("#sign-in");

  if (signinLink) {
    signinLink.addEventListener("click", () => {
      signin();
    });
  }
});
