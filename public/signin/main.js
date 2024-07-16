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
      // const RESP_GET = await fetch(API_SIGNIN_ENDPOINT)
      //   .then(resp => {
      //     switch (resp.status) {
      //       case 200:
      //         return resp;
      //       default:
      //         showError();
      //     }
      //   })
      //   .catch(err => {
      //     console.error(err);
      //     showError();
      //   });
      // if (!RESP_GET) {
      //   return;
      // }

      const PARAMS = new URLSearchParams(window.location.search);
      // if (!PARAMS.has("cb")) {
      //   showError("Invalid request.");
      //   return
      // }

      const OPTIONS = {
        method: "POST",
        headers: {
          "Content-Type": "application/json;charset=utf-8",
          //"X-CSRF-Token": RESP_GET.headers.get("x-csrf-token"),
        },
        body: JSON.stringify({
          name: nameInput.val(),
          password: passwordInput.val(),
          // client_id: PARAMS.get("client_id"),
          // redirect_uri: PARAMS.get("redirect_uri"),
          // scope: PARAMS.get("scope"),
          // state: PARAMS.get("state"),
          // nonce: PARAMS.get("nonce"),
          // response_type: PARAMS.get("response_type"),
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
