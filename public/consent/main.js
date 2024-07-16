"use strict";

const FRONT_BASE_URL = "https://localhost:4443"
const API_ENDPOINT_BASE_URL = `${FRONT_BASE_URL}/connect`;
const API_CONSENT_ENDPOINT = API_ENDPOINT_BASE_URL + "/user/consent";

function showError(msg) {
  if (!msg) {
    msg = "Some error occurred.";
  }
  $("#alert-msg").text(msg);
  $("#alert").css("display", "block");
  $("#approve-btn").prop("disabled", false);
  $("#decline-btn").prop("disabled", false);
}

$(function () {
  $("#approve-btn").click(function () {
    $("#approve-btn").prop("disabled", true);
    (async function submit() {
      const OPTIONS = {
        method: "POST",
        headers: {
          "Content-Type": "application/json;charset=utf-8",
        },
      };

      const RESP = await fetch(API_CONSENT_ENDPOINT + window.location.search, OPTIONS)
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
    })();
  });
});
