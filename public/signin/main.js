"use strict";

const FRONT_BASE_URL = "https://localhost:4443"
const API_ENDPOINT_BASE_URL = `${FRONT_BASE_URL}/connect`;
const API_SIGNIN_ENDPOINT = API_ENDPOINT_BASE_URL + "/user/authentication";

function showError(msg) {
  if (!msg) {
    msg = "Some error occurred.";
  }
  $("#alert-msg").text(msg);
  $("#alert").css("display", "block");
  $("#signin-btn").prop("disabled", false);
}

$(function () {
  const nameInput = $("#name");
  const passwordInput = $("#password");

  $("#signin-btn").click(function () {
    $("#signin-btn").prop("disabled", true);
    (async function submit() {
      const OPTIONS = {
        method: "POST",
        headers: {
          "Content-Type": "application/json;charset=utf-8",
        },
        body: JSON.stringify({
          name: nameInput.val(),
          password: passwordInput.val(),
        }),
      };

      const RESP = await fetch(API_SIGNIN_ENDPOINT + window.location.search, OPTIONS)
        .then(resp => {
          if (resp.redirected) {
            window.location.href = resp.url;
          }
          switch (resp.status) {
            case 200:
              return resp;
            case 403:
              showError(`${resp.status}: Access denied.`);
              break;
            default:
              showError(`An error occurred.`);
          }
        })
        .catch(err => {
          console.error(err.message);
          showError();
        });

      const DATA = await RESP.json();
      window.location.href = DATA["redirectUrl"];
    })();
  });

  // $("#signup-link").click(function () {
  //   window.location.href = "/signup" + window.location.search;
  // });
});
